package workflows

import (
	"context"
	"testing"

	automationService "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/automation/service"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/automation/types"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/actionlog"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/logger"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/automation"
	"github.com/stretchr/testify/require"
)

func Test_actionlog_functions(t *testing.T) {
	var (
		ctx = bypassRBAC(context.Background())
		req = require.New(t)
	)

	req.NoError(defStore.TruncateActionlogs(ctx))

	//// register action log with storage backend
	automation.ActionlogHandler(
		automationService.Registry(),
		actionlog.NewService(defStore, logger.Default(), logger.Default(), actionlog.NewPolicyAll()),
	)

	loadNewScenario(ctx, t)

	var (
		aux = struct {
			Actions actionlog.ActionSet
		}{}
	)

	vars, _ := mustExecWorkflow(ctx, t, "logger", types.WorkflowExecParams{})
	req.NoError(vars.Decode(&aux))

	// Expecting both, invoker & runner to be same as invoker
	req.Len(aux.Actions, 2)

	//undo action log registration
	automation.ActionlogHandler(
		automationService.Registry(),
		automationService.DefaultActionlog,
	)

}
