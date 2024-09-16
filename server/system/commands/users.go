package commands

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"strconv"
	"syscall"
	"time"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/auth"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/cli"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/filter"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/store"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/service"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

func Users(ctx context.Context, app serviceInitializer) *cobra.Command {
	var (
		flagNoPassword       bool
		flagPassword         string
		flagMakePasswordLink bool
		flagRoles            []string
		flagSendEmail        bool
	)

	// User management commands.
	cmd := &cobra.Command{
		Use:     "users",
		Aliases: []string{"user"},
		Short:   "User management",
	}

	// List users.
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List users",

		PreRunE: commandPreRunInitService(app),
		Run: func(cmd *cobra.Command, args []string) {
			ctx = auth.SetIdentityToContext(ctx, auth.ServiceUser())

			var (
				queryFlag = cmd.Flags().Lookup("query").Value.String()
				limitFlag = cmd.Flags().Lookup("limit").Value.String()

				limit int
				err   error
			)

			limit, err = strconv.Atoi(limitFlag)
			cli.HandleError(err)

			uf := types.UserFilter{Query: queryFlag}
			uf.Sort = filter.SortExprSet{&filter.SortExpr{Column: "updated_at"}}
			uf.Limit = uint(limit)

			users, _, err := service.DefaultStore.SearchUsers(ctx, uf)
			cli.HandleError(err)

			fmt.Fprintf(
				cmd.OutOrStdout(),
				"                     Created    Updated    EmailAddress\n",
			)

			for _, u := range users {
				upd := "---- -- --"

				if u.UpdatedAt != nil {
					upd = u.UpdatedAt.Format("2006-01-02")
				}

				fmt.Fprintf(
					cmd.OutOrStdout(),
					"%20d %s %s %-100s %s\n",
					u.ID,
					u.CreatedAt.Format("2006-01-02"),
					upd,
					u.Email,
					u.Name,
				)
			}
		},
	}

	listCmd.Flags().IntP("limit", "l", 20, "How many entry to display")
	listCmd.Flags().StringP("query", "q", "", "Query and filter by handle, email, name")

	addCmd := &cobra.Command{
		Use:   "add [email]",
		Short: "Add new user",
		Args:  cobra.MinimumNArgs(1),

		PreRunE: commandPreRunInitActivate(app),
		Run: func(cmd *cobra.Command, args []string) {
			ctx = auth.SetIdentityToContext(ctx, auth.ServiceUser())

			var (
				authSvc = service.Auth(service.AuthOptions{})

				// @todo email validation
				user = &types.User{Email: args[0]}

				err error

				// url for to create password(make-password-link)
				url string

				// Use provided password...
				password = []byte(flagPassword)

				role *types.Role
				mm   types.RoleMemberSet
			)

			if !flagNoPassword && !flagMakePasswordLink && len(password) == 0 {
				cmd.Print("Set password: ")
				if password, err = terminal.ReadPassword(int(syscall.Stdin)); err != nil {
					cli.HandleError(err)
				}
			}

			if len(password) > 0 && !authSvc.CheckPasswordStrength(string(password)) {
				cli.HandleError(service.AuthErrPasswordNotSecure())
			}

			for _, ri := range flagRoles {
				role, err = service.DefaultRole.FindByAny(ctx, ri)
				cli.HandleError(err)

				mm = append(mm, &types.RoleMember{RoleID: role.ID})
			}

			// Update current settings to be sure we do not have outdated values
			cli.HandleError(service.DefaultSettings.UpdateCurrent(ctx))

			if user, err = service.DefaultUser.Create(ctx, user); err != nil {
				cli.HandleError(err)
			}

			// generate the create password link of user
			if flagNoPassword && len(password) == 0 && flagMakePasswordLink {
				url, err = authSvc.GeneratePasswordCreateToken(ctx, user.Email)
				if err != nil {
					cli.HandleError(err)
				}
			}

			if flagSendEmail {
				if err = authSvc.SendInviteEmail(ctx, user.Email); err != nil {
					cli.HandleError(err)
				}
			}

			fmt.Fprintf(cmd.OutOrStdout(), "User created [%d].\n", user.ID)

			if flagMakePasswordLink && len(url) > 0 {
				fmt.Fprintf(cmd.OutOrStdout(), "Set password by clicking here:\n%s\n", url)
			}

			if len(mm) > 0 {
				_ = mm.Walk(func(m *types.RoleMember) error {
					m.UserID = user.ID
					return nil
				})

				cli.HandleError(store.CreateRoleMember(ctx, service.DefaultStore, mm...))
			}

			if len(password) > 0 {
				if err = authSvc.SetPassword(ctx, user.ID, string(password)); err != nil {
					cli.HandleError(err)
				}
			}
		},
	}

	addCmd.Flags().BoolVar(
		&flagNoPassword,
		"no-password",
		false,
		"Do not ask for password")

	addCmd.Flags().StringVar(
		&flagPassword,
		"password",
		"",
		"Provide password (as alternative to interactive way)")

	addCmd.Flags().BoolVar(
		&flagMakePasswordLink,
		"make-password-link",
		false,
		"Provide link to create password")

	addCmd.Flags().StringSliceVar(
		&flagRoles,
		"role",
		nil,
		"Add user to roles (use ID or handle, repeat for multiple roles)")

	addCmd.Flags().BoolVar(
		&flagSendEmail,
		"send-invite",
		false,
		"Send invite email to user with accept invite link")

	pwdCmd := &cobra.Command{
		Use:     "password [email]",
		Short:   "Change password for user",
		Args:    cobra.MinimumNArgs(1),
		PreRunE: commandPreRunInitService(app),
		Run: func(cmd *cobra.Command, args []string) {
			ctx = auth.SetIdentityToContext(ctx, auth.ServiceUser())

			var (
				user     *types.User
				err      error
				password []byte
			)

			// Update current settings to be sure that we do not have outdated values
			cli.HandleError(service.DefaultSettings.UpdateCurrent(ctx))

			if user, err = service.DefaultUser.FindByEmail(ctx, args[0]); err != nil {
				cli.HandleError(err)
			}

			cmd.Print("Set password: ")
			if password, err = terminal.ReadPassword(int(syscall.Stdin)); err != nil {
				cli.HandleError(err)
			}

			if len(password) == 0 {
				// Password not set, that's ok too.
				return
			}

			if err = service.DefaultAuth.SetPassword(ctx, user.ID, string(password)); err != nil {
				cli.HandleError(err)
			}
		},
	}

	cmd.AddCommand(
		listCmd,
		addCmd,
		pwdCmd,
		UsersSynthetic(ctx, app),
	)

	return cmd
}

func UsersSynthetic(ctx context.Context, app serviceInitializer) *cobra.Command {
	var (
		total uint
		faker = gofakeit.NewCrypto()

		synth = &cobra.Command{
			Use:     "synthetic",
			Aliases: []string{"synth"},
			Short:   "Generate or remove synthetic users",
			PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
				return app.InitServices(ctx)
			},
		}

		gen = &cobra.Command{
			Use:     "generate",
			Aliases: []string{"gen"},
			Short:   "Generate synthetic users",
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Printf("Generating %d users ...", total)
				bm := time.Now()

				ctx = auth.SetIdentityToContext(ctx, auth.ServiceUser())
				cli.HandleError(service.DefaultUser.CreateSynthetic(ctx, faker, total))

				cmd.Printf("done in %s", time.Since(bm).Round(time.Millisecond))
				cmd.Println()
			},
		}

		rem = &cobra.Command{
			Use:     "remove",
			Aliases: []string{"rm", "d", "delete", "del"},
			Short:   "Remove synthetic users",
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Printf("Removing all synthetic users ...")
				bm := time.Now()

				ctx = auth.SetIdentityToContext(ctx, auth.ServiceUser())
				cli.HandleError(service.DefaultUser.RemoveSynthetic(ctx))

				cmd.Printf("done in %s", time.Since(bm).Round(time.Millisecond))
				cmd.Println()
			},
		}
	)

	gen.Flags().UintVarP(&total, "total", "t", 1, "Number of synthetic users generated")

	synth.AddCommand(gen, rem)

	return synth
}
