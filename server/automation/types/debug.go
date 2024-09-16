package types

import (
	"context"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/expr"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/logger"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/wfexec"
	"go.uber.org/zap"
)

type (
	debugStep struct {
		wfexec.StepIdentifier
		log *zap.Logger
	}
)

// DebugStep creates a step that logs entire contents of the scope
func DebugStep(log *zap.Logger) *debugStep {
	return &debugStep{
		log: log,
	}
}

// Executes debug step
func (s debugStep) Exec(ctx context.Context, r *wfexec.ExecRequest) (wfexec.ExecResponse, error) {
	var (
		fields = make([]zap.Field, 0)

		// We want to be sure to actually get the logger that will print out debug info
		// so we need to use logger directly from workflow and not the one we (could)
		// get from context by wfexec.
		log = s.log.With(
			logger.Uint64("sessionID", r.SessionID),
			logger.Uint64("stateID", r.StateID),
			logger.Uint64("stepID", s.ID()),
		)
	)

	if r.Scope == nil {
		fields = append(fields, zap.Int("scopeKeyCount", 0))
	} else {
		fields = append(fields, zap.Int("scopeKeyCount", r.Scope.Len()))
		_ = r.Scope.Each(func(k string, v expr.TypedValue) error {
			fields = append(fields, zap.Any(k, v.Get()))
			fields = append(fields, zap.Any(k+"Type", v.Type()))
			return nil
		})
	}

	log.Debug("debug", fields...)

	return &expr.Vars{}, nil
}
