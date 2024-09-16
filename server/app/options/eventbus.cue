package options

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/codegen/schema"
)

eventbus: schema.#optionsGroup & {
	handle: "eventbus"

	imports: [
		"\"time\"",
	]

	options: {
		scheduler_enabled: {
			type:          "bool"
			defaultGoExpr: "true"
			description:   "Enable eventbus scheduler."
		}
		scheduler_interval: {
			type:        "time.Duration"
			description: "Set time interval for `eventbus` scheduler."

			defaultGoExpr: "time.Minute"
			defaultValue:  "60s"
		}
	}
	title: "Events and scheduler"
}
