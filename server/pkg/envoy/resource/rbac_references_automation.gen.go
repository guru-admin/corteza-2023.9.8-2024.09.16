package resource

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/automation/types"
)

// AutomationWorkflowRbacReferences generates RBAC references
//
// Resources with "envoy: false" are skipped
//
// This function is auto-generated
func AutomationWorkflowRbacReferences(workflow string) (res *Ref, pp []*Ref, err error) {
	if workflow != "*" {
		res = &Ref{ResourceType: types.WorkflowResourceType, Identifiers: MakeIdentifiers(workflow)}
	}

	return
}
