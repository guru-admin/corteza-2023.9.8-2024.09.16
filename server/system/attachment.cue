package system

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/codegen/schema"
)

attachment: {
	features: {
		labels: false
	}

	model: {
		attributes: {
			id: schema.IdField
			owner_id:   {
				storeIdent: "rel_owner",
				ident: "ownerID"
				schema.AttributeUserRef,
			}
			kind: {
				sortable: true
				dal: {}
			}
			url: {
				dal: {}
			}
			preview_url: {
				dal: {}
			}
			name: {
				sortable: true
				dal: {}
			}
			meta: {
				goType: "types.AttachmentMeta"
				dal: { type: "JSON", defaultEmptyObject: true }
				omitSetter: true
				omitGetter: true
			}
			created_at: schema.SortableTimestampNowField
			updated_at: schema.SortableTimestampNilField
			deleted_at: schema.SortableTimestampNilField
		}

		indexes: {
			"primary": { attribute: "id" }
		}
	}

	envoy: {
		omit: true
	}

	filter: {
		struct: {
			kind: {}
		}

		byValue: ["kind"]
	}

	store: {
		api: {
			lookups: [
				{ fields: ["id"] },
			]
		}
	}
}
