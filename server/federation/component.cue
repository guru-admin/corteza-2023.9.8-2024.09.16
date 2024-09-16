package federation

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/codegen/schema"
)

component: schema.#component & {
	handle: "federation"

	resources: {
		"node":           node
		"node-sync":      nodeSync
		"exposed-module": exposedModule
		"shared-module":  sharedModule
		"module-mapping": moduleMapping
	}

	envoy: {
		omit: true
	}

	rbac: operations: {
		"grant": description:           "Manage federation permissions"
		"pair": description:            "Pair federation nodes"
		"settings.read": description:   "Read settings"
		"settings.manage": description: "Manage settings"
		"node.create": description:     "Create new federation node"
		"nodes.search": description:    "List, search or filter federation nodes"
	}
}
