package automation

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/codegen/schema"
)

component: schema.#component & {
	handle: "automation"

	resources: {
		"workflow": workflow
		"session":  session
		"trigger":  trigger
	}

	rbac: operations: {
		"grant": description:                        "Manage automation permissions"
		"workflow.create": description:              "Create workflows"
		"triggers.search": description:              "List, search or filter triggers"
		"sessions.search": description:              "List, search or filter sessions"
		"workflows.search": description:             "List, search or filter workflows"
		"resource-translations.manage": description: "List, search, create, or update resource translations"
	}
}
