package rest

import (
	"context"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/automation/rest/request"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/automation/service"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/automation/types"
)

type (
	Function struct {
		reg interface {
			Functions() []*types.Function
		}
	}

	functionSetPayload struct {
		Set []*types.Function `json:"set"`
	}
)

func (Function) New() *Function {
	ctrl := &Function{reg: service.Registry()}
	return ctrl
}

func (ctrl Function) List(_ context.Context, _ *request.FunctionList) (interface{}, error) {
	return functionSetPayload{Set: ctrl.reg.Functions()}, nil
}
