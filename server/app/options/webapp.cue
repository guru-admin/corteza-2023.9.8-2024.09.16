package options

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/codegen/schema"
)

webapp: schema.#optionsGroup & {
	handle: "webapp"

	options: {
		scss_dir_path: {
			description:  "Path to custom SCSS source files directory"
		}
	}
}
