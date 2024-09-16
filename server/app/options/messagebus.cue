package options

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/codegen/schema"
)

messagebus: schema.#optionsGroup & {
	handle: "messagebus"
	options: {
		Enabled: {
			type:          "bool"
			defaultGoExpr: "true"
			description:   "Enable messagebus"
		}
		log_enabled: {
			type:        "bool"
			description: "Enable extra logging for messagebus watchers"
		}
	}
	title: "Messaging queue"
}
