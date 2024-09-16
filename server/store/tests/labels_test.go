package tests

import (
	"context"
	"testing"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/label/types"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/logger"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/store"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/require"
)

func testLabels(t *testing.T, s store.Labels) {
	var (
		ctx = context.Background()
	)

	ctx = logger.ContextWithValue(ctx, logger.MakeDebugLogger())

	t.Run("create", func(t *testing.T) {
		req := require.New(t)
		req.NoError(s.TruncateLabels(ctx))
		req.NoError(s.CreateLabel(ctx, &types.Label{
			Kind:       "kind",
			ResourceID: 1,
			Name:       "lname",
			Value:      "lvalue",
		}))
	})

	t.Run("update", func(t *testing.T) {
		req := require.New(t)
		req.NoError(s.TruncateLabels(ctx))
		req.NoError(s.UpdateLabel(ctx, &types.Label{
			Kind:       "kind",
			ResourceID: 1,
			Name:       "lname",
			Value:      "lvalue",
		}))
	})

	t.Run("upsert", func(t *testing.T) {
		req := require.New(t)
		req.NoError(s.TruncateLabels(ctx))
		req.NoError(s.UpsertLabel(ctx, &types.Label{
			Kind:       "kind",
			ResourceID: 1,
			Name:       "lname",
			Value:      "lvalue",
		}))
	})
}
