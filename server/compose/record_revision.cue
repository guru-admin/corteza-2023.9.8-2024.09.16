package compose

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/codegen/schema"
)

record_revision: {
	model: {
		ident: "compose_record_revisions"
		omitGetterSetter: true

		attributes: {
			id: schema.IdField
			timestamp: schema.SortableTimestampField & { storeIdent: "ts" }
			rel_resource: {
			 	ident: "resourceID",
				goType: "uint64",
				dal: { type: "ID" }
			}
			revision: {
				goType: "uint"
				dal: { type: "Number", meta: { "rdbms:type": "integer" } }
			}
			operation: {
				dal: {}
			}
			rel_user:   schema.AttributeUserRef
			delta: {
				goType: "types.RecordValueSet",
				dal: { type: "JSON", defaultEmptyObject: true }
			}
			comment: {
				dal: {}
			}
		}

		indexes: {
			"primary": { attribute: "id" }
		}
	}

	envoy: {
		omit: true
	}
}
