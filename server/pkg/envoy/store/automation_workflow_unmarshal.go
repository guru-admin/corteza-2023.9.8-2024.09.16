package store

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/automation/types"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy/resource"
)

func newAutomationWorkflow(wf *types.Workflow, tt types.TriggerSet, ux *userIndex) *automationWorkflow {
	return &automationWorkflow{
		wf: wf,
		tt: tt,

		ux: ux,
	}
}

func (awf *automationWorkflow) MarshalEnvoy() ([]resource.Interface, error) {
	rs := resource.NewAutomationWorkflow(awf.wf)
	syncUserStamps(rs.Userstamps(), awf.ux)

	for _, t := range awf.tt {
		rt := rs.AddAutomationTrigger(t)
		syncUserStamps(rt.Userstamps(), awf.ux)
	}

	for _, s := range awf.wf.Steps {
		rs.AddAutomationWorkflowStep(s)
	}

	for _, p := range awf.wf.Paths {
		rs.AddAutomationWorkflowPath(p)
	}

	return envoy.CollectNodes(
		rs,
	)
}
