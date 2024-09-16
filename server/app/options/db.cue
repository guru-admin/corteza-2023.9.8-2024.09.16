package options

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/codegen/schema"
)

DB: schema.#optionsGroup & {
	handle: "DB"
	options: {
		DSN: {
			defaultValue: "sqlite3://file::memory:?cache=shared&mode=memory"
			description:  "Database connection string."
		}
		
	}
	title: "Connection to data store backend"
}
