package workflows

import (
	"context"
	"fmt"
	"testing"

	autTypes "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/automation/types"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/expr"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/wfexec"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/automation"
	"github.com/stretchr/testify/require"
)

func Test0006_iterator_users(t *testing.T) {
	wfexec.MaxIteratorBufferSize = wfexec.DefaultMaxIteratorBufferSize
	defer func() {
		wfexec.MaxIteratorBufferSize = wfexec.DefaultMaxIteratorBufferSize
	}()

	var (
		ctx = bypassRBAC(context.Background())
		req = require.New(t)
	)

	req.NoError(defStore.TruncateUsers(ctx))

	loadScenario(ctx, t)

	var (
		_, trace = mustExecWorkflow(ctx, t, "testing", autTypes.WorkflowExecParams{})
	)

	// 6x iterator, 5x continue, 1x terminator, 1x completed
	req.Len(trace, 13)

	// there are 4 iterator calls; each on the *2 index
	ctr := int64(-1)
	for j := 0; j <= 4; j++ {
		ix := j * 2
		ctr++

		frame := trace[ix]
		req.Equal(uint64(10), frame.StepID)

		i, err := expr.Integer{}.Cast(frame.Results.GetValue()["i"])
		req.NoError(err)
		req.Equal(ctr, i.Get().(int64))

		usr, err := automation.NewUser(frame.Results.GetValue()["u"])
		req.NoError(err)
		req.Equal(fmt.Sprintf("u%d", ctr+1), usr.GetValue().Handle)
	}
}

func Test0006_iterator_users_chunked(t *testing.T) {
	wfexec.MaxIteratorBufferSize = 2
	defer func() {
		wfexec.MaxIteratorBufferSize = wfexec.DefaultMaxIteratorBufferSize
	}()

	var (
		ctx = bypassRBAC(context.Background())
		req = require.New(t)
	)

	req.NoError(defStore.TruncateUsers(ctx))

	loadScenarioWithName(ctx, t, "S0006_iterator_users")

	var (
		_, trace = mustExecWorkflow(ctx, t, "testing", autTypes.WorkflowExecParams{})
	)

	// 6x iterator, 5x continue, 1x terminator, 1x completed
	req.Len(trace, 13)

	// there are 4 iterator calls; each on the *2 index
	ctr := int64(-1)
	for j := 0; j <= 4; j++ {
		ix := j * 2
		ctr++

		frame := trace[ix]
		req.Equal(uint64(10), frame.StepID)

		i, err := expr.Integer{}.Cast(frame.Results.GetValue()["i"])
		req.NoError(err)
		req.Equal(ctr, i.Get().(int64))

		usr, err := automation.NewUser(frame.Results.GetValue()["u"])
		req.NoError(err)
		req.Equal(fmt.Sprintf("u%d", ctr+1), usr.GetValue().Handle)
	}
}

func Test0006_iterator_users_limited(t *testing.T) {
	var (
		ctx = bypassRBAC(context.Background())
		req = require.New(t)
	)

	req.NoError(defStore.TruncateUsers(ctx))

	loadScenarioWithName(ctx, t, "iterator_users_limit")

	var (
		_, trace = mustExecWorkflow(ctx, t, "testing", autTypes.WorkflowExecParams{})
	)

	// 3x iterator, 2x continue, 1x terminator, 1x completed
	req.Len(trace, 7)

	// there are 4 iterator calls; each on the *2 index
	ctr := int64(-1)
	for j := 0; j <= 1; j++ {
		ix := j * 2
		ctr++

		frame := trace[ix]
		req.Equal(uint64(10), frame.StepID)

		i, err := expr.Integer{}.Cast(frame.Results.GetValue()["i"])
		req.NoError(err)
		req.Equal(ctr, i.Get().(int64))

		usr, err := automation.NewUser(frame.Results.GetValue()["u"])
		req.NoError(err)
		req.Equal(fmt.Sprintf("u%d", ctr+1), usr.GetValue().Handle)
	}
}
