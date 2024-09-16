package yaml

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/automation/types"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy/resource"
)

type (
	automationWorkflow struct {
		res      *types.Workflow
		triggers automationTriggerSet
		steps    automationWorkflowStepSet
		paths    automationWorkflowPathSet

		ts *resource.Timestamps
		us *resource.Userstamps

		envoyConfig   *resource.EnvoyConfig
		encoderConfig *EncoderConfig

		rbac rbacRuleSet
	}
	automationWorkflowSet []*automationWorkflow

	automationTrigger struct {
		res *types.Trigger

		ts *resource.Timestamps
		us *resource.Userstamps

		envoyConfig   *resource.EnvoyConfig
		encoderConfig *EncoderConfig
	}
	automationTriggerSet []*automationTrigger

	automationWorkflowStep struct {
		res *types.WorkflowStep

		ts *resource.Timestamps
		us *resource.Userstamps

		envoyConfig   *resource.EnvoyConfig
		encoderConfig *EncoderConfig
	}
	automationWorkflowStepSet []*automationWorkflowStep

	automationWorkflowPath struct {
		res *types.WorkflowPath

		ts *resource.Timestamps
		us *resource.Userstamps

		envoyConfig   *resource.EnvoyConfig
		encoderConfig *EncoderConfig
	}
	automationWorkflowPathSet []*automationWorkflowPath
)

func (nn automationWorkflowSet) configureEncoder(cfg *EncoderConfig) {
	for _, n := range nn {
		n.encoderConfig = cfg
	}
}
