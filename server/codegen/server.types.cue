package codegen

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/app"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/codegen/schema"
)

[...schema.#codegen] &
[
	for cmp in app.corteza.components {
		template: "gocode/types/$component_resources.go.tpl"
		output:   "\(cmp.ident)/types/resources.gen.go"
		payload: {
			package: "types"

			cmpIdent: cmp.ident
			// Operation/resource validators, grouped by resource
			types: [
				for res in cmp.resources {
					const:   "\(res.expIdent)ResourceType"
					type:    res.fqrt
				},
				{
					const:     "ComponentResourceType"
					type:      cmp.fqrt
				},
			]
		}
	},
	for cmp in app.corteza.components {
		template: "gocode/types/$component_getters_setters.go.tpl"
		output:   "\(cmp.ident)/types/getters_setters.gen.go"
		payload: {
			package: "types"

			cmpIdent: cmp.ident
			resources: [ for res in cmp.resources { res }]
		}
	},
]
