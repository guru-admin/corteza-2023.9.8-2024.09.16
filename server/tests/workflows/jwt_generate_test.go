package workflows

import (
	"context"
	_ "embed"
	"strings"
	"testing"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/automation/types"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/jwt_generate/match_token
var tkn string

func Test_jwt_generate(t *testing.T) {
	var (
		ctx = bypassRBAC(context.Background())
		req = require.New(t)

		aux = struct {
			Token string
		}{}
	)

	req.NoError(defStore.TruncateAttachments(ctx))

	loadNewScenario(ctx, t)

	vars, _ := mustExecWorkflow(ctx, t, "jwt_generate", types.WorkflowExecParams{})

	req.NoError(vars.Decode(&aux))
	req.Equal(strings.TrimSuffix(tkn, "\n"), aux.Token)
}
