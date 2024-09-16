package service

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/automation/types"
	sysEvent "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/service/event"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidateWorkflowTriggersEmpty(t *testing.T) {
	var (
		req = require.New(t)

		issues = validateWorkflowTriggers(
			&types.Workflow{},
		)
	)

	req.Empty(issues)
}
func TestValidateWorkflowTriggersRunAs(t *testing.T) {
	var (
		req = require.New(t)
		soi = sysEvent.SystemOnInterval()

		issues = validateWorkflowTriggers(
			&types.Workflow{},
			&types.Trigger{
				Enabled:      true,
				ResourceType: soi.ResourceType(),
				EventType:    soi.EventType(),
			},
		)
	)

	req.Len(issues, 1)
	req.Contains(issues[0].String(), "requires run-as to be set")
}

func TestValidateWorkflowTriggersSubWorkflow(t *testing.T) {
	var (
		req = require.New(t)

		issues = validateWorkflowTriggers(
			&types.Workflow{Meta: &types.WorkflowMeta{SubWorkflow: true}},
			&types.Trigger{Enabled: true},
		)
	)

	req.Len(issues, 1)
	req.Contains(issues[0].String(), "marked as sub-workflow")
}
