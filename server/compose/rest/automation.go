package rest

import (
	"context"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/api"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/compose/rest/request"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/compose/service/event"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/corredor"
)

type (
	Automation struct{}
)

func (Automation) New() *Automation {
	return &Automation{}
}

func (ctrl *Automation) List(ctx context.Context, r *request.AutomationList) (interface{}, error) {
	return corredor.GenericListHandler(
		ctx,
		corredor.Service(),
		corredor.Filter{
			ResourceTypePrefixes: r.ResourceTypePrefixes,
			ExcludeInvalid:       r.ExcludeInvalid,
			ResourceTypes:        r.ResourceTypes,
			EventTypes:           r.EventTypes,
			ExcludeServerScripts: r.ExcludeServerScripts,
			ExcludeClientScripts: r.ExcludeClientScripts,
		},
		"compose",
	)
}

func (ctrl *Automation) Bundle(ctx context.Context, r *request.AutomationBundle) (interface{}, error) {
	return corredor.GenericBundleHandler(
		ctx,
		corredor.Service(),
		r.Bundle,
		r.Type,
		r.Ext,
	)
}

func (ctrl *Automation) TriggerScript(ctx context.Context, r *request.AutomationTriggerScript) (interface{}, error) {
	return api.OK(), corredor.Service().Exec(ctx, r.Script, corredor.ExtendScriptArgs(event.ComposeOnManual(), r.Args))
}
