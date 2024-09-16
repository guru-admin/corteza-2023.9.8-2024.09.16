package options

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/codegen/schema"
)

limit: schema.#optionsGroup & {
	handle: "limit"
	options: {
		system_users: {
			type:        "int"
			description: "Maximum number of valid (not deleted, not suspended) users"
		}
	}
	title: "Limits"
}
