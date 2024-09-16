package envoy

import (
	"context"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/compose/dalutils"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/compose/types"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/dal"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoyx"
	"github.com/spf13/cast"
)

type (
	// recordGetter is a utility struct to resolve record references from
	// different parts of the system such as the dep graph and the database
	recordGetter struct {
		// Bits to get data from the dep graph
		relDatasource *RecordDatasource

		// Bits to get data from the database
		dalSvc     dal.FullService
		relModule  types.Module
		baseFilter types.RecordFilter
	}
)

func makeRecordGetter(dalSvc dal.FullService, tt envoyx.Traverser, n *envoyx.Node, modRef, dsRef envoyx.Ref) (g *recordGetter) {
	g = &recordGetter{
		dalSvc: dalSvc,
	}

	// Resolve from dep graph
	auxDs := tt.ParentForRef(n, dsRef)
	if auxDs != nil {
		g.relDatasource = auxDs.Datasource.(*RecordDatasource)
	}

	// Resolve from the database
	mod := g.getRefMod(tt, n, modRef)
	if mod != nil {
		g.baseFilter = types.RecordFilter{
			ModuleID:    mod.ID,
			NamespaceID: mod.NamespaceID,
		}
	}

	return
}

// resolve resolves the provided reference into a record ID; 0 if can't be resolved
func (g *recordGetter) resolve(ctx context.Context, ref any) (out uint64, err error) {
	// Try to get from datasource
	if g.relDatasource != nil {
		out, err = g.getDS(ref)
		if err != nil {
			return
		}
	}
	if out > 0 {
		return
	}

	// Fallback to the store
	out, err = g.getDB(ctx, ref)
	if err != nil {
		return
	}
	if out > 0 {
		return
	}
	return
}

func (g *recordGetter) getDS(ref any) (out uint64, err error) {
	return g.relDatasource.ResolveRef(ref)
}

func (g *recordGetter) getDB(ctx context.Context, ref any) (out uint64, err error) {
	// @note the old version only resolved IDs so that's what we're doing here also
	// @todo consider expanding this
	id, err := cast.ToUint64E(ref)
	if err != nil {
		return 0, nil
	}

	aux, err := dalutils.ComposeRecordsFind(ctx, g.dalSvc, &g.relModule, id)
	if err != nil {
		return
	}

	out = aux.ID
	return
}

func (g *recordGetter) getRefMod(tt envoyx.Traverser, n *envoyx.Node, ref envoyx.Ref) (refMod *types.Module) {
	aux := tt.ParentForRef(n, ref)
	if aux == nil {
		return nil
	}

	return aux.Resource.(*types.Module)
}
