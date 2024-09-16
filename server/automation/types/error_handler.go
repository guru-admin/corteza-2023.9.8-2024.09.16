package types

import (
	"context"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/expr"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/wfexec"
)

type (
	errorHandlerStep struct {
		wfexec.StepIdentifier
		handler wfexec.Step
		results *expr.Vars
	}
)

func ErrorHandlerStep(h wfexec.Step, rr *expr.Vars) *errorHandlerStep {
	return &errorHandlerStep{handler: h, results: rr}
}

// Exec errorHandler step
func (h errorHandlerStep) Exec(ctx context.Context, r *wfexec.ExecRequest) (wfexec.ExecResponse, error) {
	return wfexec.ErrorHandler(h.handler, h.results), nil
}
