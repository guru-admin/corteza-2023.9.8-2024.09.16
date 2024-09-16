package documents

import (
	"context"
	"fmt"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/discovery/service"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/errors"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/id"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/options"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/filter"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/rbac"
	sysService "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/service"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
)

type (
	systemResources struct {
		opt      options.DiscoveryOpt
		settings *types.AppSettings

		rbac interface {
			SignificantRoles(res rbac.Resource, op string) (aRR, dRR []uint64)
		}

		ac interface {
			CanReadUser(ctx context.Context, r *types.User) bool
		}
		usr interface {
			Find(context.Context, types.UserFilter) (types.UserSet, types.UserFilter, error)
		}
	}
)

func SystemResources() *systemResources {
	return &systemResources{
		opt:      service.DefaultOption,
		settings: sysService.CurrentSettings,
		rbac:     rbac.Global(),
		ac:       sysService.DefaultAccessControl,
		usr:      sysService.DefaultUser,
	}
}

func (d systemResources) Users(ctx context.Context, limit uint, cur string, userID uint64, deleted uint) (rsp *Response, err error) {
	return rsp, func() (err error) {
		if !d.settings.Discovery.SystemUsers.Enabled {
			return errors.Internal("system user indexing disabled")
		}

		var (
			uu types.UserSet
			f  = types.UserFilter{
				Deleted: filter.State(deleted),
			}
		)

		if userID > 0 {
			f.UserID = append(f.UserID, id.String(userID))
		}

		if f.Paging, err = filter.NewPaging(limit, cur); err != nil {
			return err
		}

		if uu, f, err = d.usr.Find(ctx, f); err != nil {
			return err
		}

		rsp = &Response{
			Documents: make([]Document, len(uu)),
			Filter: Filter{
				Limit:    limit,
				NextPage: f.NextPage,
			},
		}

		for i, u := range uu {
			doc := &docUser{
				ResourceType: "system:user",
				UserID:       u.ID,
				Email:        u.Email,
				Name:         u.Name,
				Handle:       u.Handle,
				Suspended:    u.SuspendedAt,
				Created:      makePartialChange(&u.CreatedAt),
				Updated:      makePartialChange(u.UpdatedAt),
				Deleted:      makePartialChange(u.DeletedAt),
			}
			if len(d.opt.CortezaDomain) > 0 && u.ID > 0 {
				doc.Url = fmt.Sprintf("%s/admin/system/user/edit/%d", d.opt.CortezaDomain, u.ID)
			}

			doc.Security.AllowedRoles, doc.Security.DeniedRoles = d.rbac.SignificantRoles(u, "read")

			rsp.Documents[i].ID = u.ID
			rsp.Documents[i].Source = doc
		}

		return nil
	}()
}
