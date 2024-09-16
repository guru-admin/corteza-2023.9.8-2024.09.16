package workflows

import (
	"context"
	"testing"

	autTypes "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/automation/types"
	"github.com/stretchr/testify/require"
)

func Test0013_gateway_exclusive(t *testing.T) {
	var (
		ctx = bypassRBAC(context.Background())
		req = require.New(t)
	)

	loadScenario(ctx, t)

	t.Run("first path match", func(t *testing.T) {
		_, trace := mustExecWorkflow(ctx, t, "case1", autTypes.WorkflowExecParams{})

		req.Len(trace, 4)
		req.Equal(uint64(10), trace[0].StepID)
		req.Equal(uint64(11), trace[1].StepID)
		req.Equal(uint64(14), trace[2].StepID)
		req.Equal(uint64(0), trace[3].StepID)
	})

	t.Run("second path match", func(t *testing.T) {
		_, trace := mustExecWorkflow(ctx, t, "case2", autTypes.WorkflowExecParams{})

		req.Len(trace, 4)
		req.Equal(uint64(10), trace[0].StepID)
		req.Equal(uint64(12), trace[1].StepID)
		req.Equal(uint64(14), trace[2].StepID)
		req.Equal(uint64(0), trace[3].StepID)
	})

	t.Run("default path match", func(t *testing.T) {
		_, trace := mustExecWorkflow(ctx, t, "case3", autTypes.WorkflowExecParams{})

		req.Len(trace, 4)
		req.Equal(uint64(10), trace[0].StepID)
		req.Equal(uint64(13), trace[1].StepID)
		req.Equal(uint64(14), trace[2].StepID)
		req.Equal(uint64(0), trace[3].StepID)
	})
}
