package options

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/codegen/schema"
)

discovery: schema.#optionsGroup & {
	handle: "discovery"
	options: {
		enabled: {
			type:          "bool"
			defaultGoExpr: "false"
			description:   "Enable discovery endpoints"
		}
		debug: {
			type:          "bool"
			defaultGoExpr: "false"
			description:   "Enable discovery related activity info"
		}
		corteza_domain: {
			type:        "string"
			description: "Indicates host of corteza compose webapp"
		}
		base_url: {
			type:        "string"
			description: "Indicates host of corteza discovery server"
		}

	}
	title: "Discovery"
}
