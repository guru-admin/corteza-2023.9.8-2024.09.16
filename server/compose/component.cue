package compose

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/codegen/schema"
)

component: schema.#component & {
	handle: "compose"

	resources: {
		"attachment":          attachment
		"chart":               chart
		"module":              module
		"module-field":        moduleField
		"namespace":           namespace
		"page":                page
		"page-layout":         pageLayout
		"record":              record
		"record-revision":     record_revision
	}

	rbac: operations: {
		"settings.read": description:                "Read settings"
		"settings.manage": description:              "Manage settings"
		"namespace.create": description:             "Create namespace"
		"namespaces.search": description:            "List, search or filter namespaces"
		"resource-translations.manage": description: "List, search, create, or update resource translations"
	}
}

