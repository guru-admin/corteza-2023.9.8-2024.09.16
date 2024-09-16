package options

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/codegen/schema"
)

actionLog: schema.#optionsGroup & {
	handle: "action-log"
	env: "ACTIONLOG"
	options: {
		enabled: {
			type:          "bool"
			defaultGoExpr: "true"
		}
		debug: {
			type: "bool"
		}
		workflow_functions_enabled: {
			type: "bool"
		}
	}
	title: "Actionlog"
}
