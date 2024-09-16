package options

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/codegen/schema"
)

environment: schema.#optionsGroup & {
	handle: "environment"
	options: {
		environment: {
			defaultValue: "production"
			env:          "ENVIRONMENT"
		}
	}
	title: "Environment"
}
