package tests

import (
	"context"
	"testing"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/flag/types"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/store"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/require"
)

func testFlags(t *testing.T, s store.Flags) {
	var (
		ctx = context.Background()
	)

	t.Run("create", func(t *testing.T) {
		req := require.New(t)
		req.NoError(s.TruncateFlags(ctx))
		req.NoError(s.CreateFlag(ctx, &types.Flag{
			Kind:       "kind",
			ResourceID: 1,
			OwnedBy:    2,
			Name:       "fname",
			Active:     true,
		}))
	})

	t.Run("update", func(t *testing.T) {
		req := require.New(t)
		req.NoError(s.TruncateFlags(ctx))
		req.NoError(s.UpdateFlag(ctx, &types.Flag{
			Kind:       "kind",
			ResourceID: 1,
			OwnedBy:    2,
			Name:       "fname",
			Active:     false,
		}))
	})

	t.Run("upsert", func(t *testing.T) {
		req := require.New(t)
		req.NoError(s.TruncateFlags(ctx))
		req.NoError(s.UpsertFlag(ctx, &types.Flag{
			Kind:       "kind",
			ResourceID: 1,
			OwnedBy:    2,
			Name:       "fname",
			Active:     true,
		}))
	})
}
