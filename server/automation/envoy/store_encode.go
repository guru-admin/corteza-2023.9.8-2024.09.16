package envoy

import (
	"context"
	"time"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/automation/types"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoyx"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/store"
)

func (e StoreEncoder) encode(ctx context.Context, p envoyx.EncodeParams, s store.Storer, rt string, nn envoyx.NodeSet, tree envoyx.Traverser) (err error) {
	return
}

func (e StoreEncoder) setWorkflowDefaults(res *types.Workflow) (err error) {
	if res.CreatedAt.IsZero() {
		res.CreatedAt = time.Now()
	}
	return
}

func (e StoreEncoder) validateWorkflow(res *types.Workflow) (err error) {
	return
}

func (e StoreEncoder) setTriggerDefaults(res *types.Trigger) (err error) {
	if res.CreatedAt.IsZero() {
		res.CreatedAt = time.Now()
	}
	return
}

func (e StoreEncoder) validateTrigger(res *types.Trigger) (err error) {
	return
}
