package ctx

import (
	"context"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/apigw/profiler"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/apigw/types"
)

type (
	scopeCtxKey    struct{}
	profilerCtxKey struct{}
)

func ScopeToContext(ctx context.Context, s *types.Scp) context.Context {
	return context.WithValue(ctx, scopeCtxKey{}, s)
}

func ScopeFromContext(ctx context.Context) (ss *types.Scp) {
	s := ctx.Value(scopeCtxKey{})

	if s == nil {
		return &types.Scp{}
	}

	return s.(*types.Scp)
}

func ProfilerToContext(ctx context.Context, h interface{}) context.Context {
	return context.WithValue(ctx, profilerCtxKey{}, h)
}

func ProfilerFromContext(ctx context.Context) (h *profiler.Hit) {
	hh := ctx.Value(profilerCtxKey{})

	if hh == nil {
		return nil
	}

	return hh.(*profiler.Hit)
}
