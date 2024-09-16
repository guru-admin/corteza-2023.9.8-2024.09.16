package workflows

import (
	"context"
	"testing"

	autTypes "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/automation/types"
	sysTypes "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
	"github.com/stretchr/testify/require"
)

func Test0003_create_user(t *testing.T) {
	var (
		ctx = bypassRBAC(context.Background())
		req = require.New(t)
	)

	loadScenario(ctx, t)

	var (
		aux = struct {
			User *sysTypes.User
		}{}
		vars, _ = mustExecWorkflow(ctx, t, "create-user", autTypes.WorkflowExecParams{})
	)

	req.NoError(vars.Decode(&aux))
}
