package service

import (
	"context"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/compose/types"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/store"
	sysTypes "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func resourceMaker(ctx context.Context, t *testing.T, s store.Storer, mods ...any) {
	req := require.New(t)

	for _, m := range mods {
		switch c := m.(type) {
		case *sysTypes.User:
			t.Log("creating User")
			req.NoError(store.CreateUser(ctx, s, c))

		case *sysTypes.Role:
			t.Log("creating Role")
			req.NoError(store.CreateRole(ctx, s, c))

		case *types.Namespace:
			t.Log("creating Namespace")
			req.NoError(store.CreateComposeNamespace(ctx, s, c))

		case *types.Module:
			t.Log("creating Module")
			req.NoError(store.CreateComposeModule(ctx, s, c))

		case *types.ModuleField:
			t.Log("creating ModuleField")
			req.NoError(store.CreateComposeModuleField(ctx, s, c))
		}
	}
}
