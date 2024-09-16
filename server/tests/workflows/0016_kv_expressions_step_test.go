package workflows

import (
	"context"
	cmpTypes "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/compose/types"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/expr"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/automation/types"
)

func Test0016_kv_expressions_step(t *testing.T) {
	var (
		ctx = bypassRBAC(context.Background())
		req = require.New(t)
	)

	loadScenario(ctx, t)

	t.Run("KV related expressions", func(t *testing.T) {
		type (
			testInput struct {
				Out map[string]string
			}
		)
		var (
			aux      = testInput{}
			vars, _  = mustExecWorkflow(ctx, t, "kv_expressions", types.WorkflowExecParams{})
			expected = testInput{
				Out: map[string]string{
					"testString": "testing string",
				},
			}
		)

		req.NoError(vars.Decode(&aux))
		req.Equal(expected, aux)
	})

	t.Run("KVV related expressions", func(t *testing.T) {
		type (
			testInput struct {
				Out map[string][]string
			}
		)
		var (
			aux      = testInput{}
			vars, _  = mustExecWorkflow(ctx, t, "kvv_expressions", types.WorkflowExecParams{})
			expected = testInput{
				Out: map[string][]string{
					"testString": {"foo", "bar"},
				},
			}
		)

		req.NoError(vars.Decode(&aux))
		req.Equal(expected, aux)
	})

	t.Run("Vars related expressions", func(t *testing.T) {
		type (
			testInput struct {
				Out map[string]expr.TypedValue
			}
		)
		var (
			aux     = testInput{}
			vars, _ = mustExecWorkflow(ctx, t, "vars_expressions", types.WorkflowExecParams{})

			testString = expr.Must(expr.NewString("testing string"))
			testInt    = expr.Must(expr.NewInteger(40))
			testVar    = expr.Must(expr.NewVars(map[string]interface{}{
				"testString": testString,
				"testFloat":  expr.Must(expr.NewFloat(50)),
			}))
			expectedVars = map[string]expr.TypedValue{
				"testString": testString,
				"testInt":    testInt,
				"testVar":    testVar,
			}

			expected = testInput{
				Out: expectedVars,
			}
		)

		req.NoError(vars.Decode(&aux))
		req.Equal(expected, aux)
	})

	t.Run("ComposeRecordValues related expressions", func(t *testing.T) {
		type (
			testInput struct {
				Out *cmpTypes.Record
			}
		)
		var (
			aux      = testInput{}
			vars, _  = mustExecWorkflow(ctx, t, "compose_record_values_expressions", types.WorkflowExecParams{})
			expected = testInput{
				Out: &cmpTypes.Record{
					Values: []*cmpTypes.RecordValue{
						{
							Name:  "testFloat",
							Value: "50",
						},
					},
				},
			}
		)

		req.NoError(vars.Decode(&aux))
		req.Equal(expected, aux)
	})
}
