package commands

import (
	"context"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/cli"
	"github.com/spf13/cobra"
)

type (
	serviceInitializer interface {
		InitServices(ctx context.Context) error
		Activate(ctx context.Context) error
	}
)

func commandPreRunInitService(app serviceInitializer) func(*cobra.Command, []string) error {
	return func(_ *cobra.Command, _ []string) error {
		return app.InitServices(cli.Context())
	}
}

func commandPreRunInitActivate(app serviceInitializer) func(*cobra.Command, []string) error {
	return func(_ *cobra.Command, _ []string) error {
		return app.Activate(cli.Context())
	}
}
