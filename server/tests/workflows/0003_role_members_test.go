package workflows

import (
	"context"
	"testing"

	autTypes "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/automation/types"
	sysTypes "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
	"github.com/stretchr/testify/require"
)

func Test0003_role_members(t *testing.T) {
	var (
		ctx = bypassRBAC(context.Background())
		req = require.New(t)
	)

	req.NoError(defStore.TruncateRoleMembers(ctx))
	req.NoError(defStore.TruncateUsers(ctx))
	req.NoError(defStore.TruncateRoles(ctx))

	loadScenario(ctx, t)

	var (
		aux = struct {
			UsersPre_r1 sysTypes.UserSet
			TotalPre_r1 uint64
			UsersPre_r2 sysTypes.UserSet
			TotalPre_r2 uint64

			UsersPost_r1 sysTypes.UserSet
			TotalPost_r1 uint64
			UsersPost_r2 sysTypes.UserSet
			TotalPost_r2 uint64
		}{}
		vars, _ = mustExecWorkflow(ctx, t, "role-members", autTypes.WorkflowExecParams{})
	)

	req.NoError(vars.Decode(&aux))

	req.Len(aux.UsersPre_r1, 1)
	req.Equal(uint64(1), aux.TotalPre_r1)
	req.Len(aux.UsersPre_r2, 1)
	req.Equal(uint64(1), aux.TotalPre_r2)

	req.Len(aux.UsersPost_r1, 0)
	req.Equal(uint64(0), aux.TotalPost_r1)
	req.Len(aux.UsersPost_r2, 1)
	req.Equal(uint64(1), aux.TotalPost_r2)
}
