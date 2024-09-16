package codegen

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/app"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/codegen/schema"
)

[...schema.#codegen] &
[
	{
		template: "docs/options.adoc.tpl"
		output:   "src/modules/generated/partials/env-options.gen.adoc"
		payload: {
			groups: [
				for g in app.corteza.options {
					title: g.title
					intro?: g.intro

					options: g.options
				},
			]
		}
	},
]
