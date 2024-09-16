package system

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/codegen/schema"
)

apigw_filter: {
	features: {
		labels: false
	}

	model: {
		attributes: {
			id: schema.IdField
			route:  {
				sortable: true, goType: "uint64", storeIdent: "rel_route"
				dal: { type: "Ref", refModelResType: "corteza::system:apigw-route" }
				identAlias: ["route", "Route", "ApigwRouteID"]
				envoy: {
					store: {
						omitRefFilter: true
					}
				}
		  }
			weight: {
			  sortable: true,
			  goType: "uint64"
			  dal: { type: "Number", meta: { "rdbms:type": "integer" } }
			}
			kind: {
				sortable: true
				dal: { type: "Text", length: 64 }
			}
			ref: {
				dal: { type: "Text", length: 64 }
			}
			enabled: {
				sortable: true,
				goType: "bool"
				dal: { type: "Boolean" }
			}
			params: {
				goType: "types.ApigwFilterParams"
				dal: { type: "JSON", defaultEmptyObject: true }
				omitSetter: true
				omitGetter: true
			}

			created_at: schema.SortableTimestampNowField
			updated_at: schema.SortableTimestampNilField
			deleted_at: schema.SortableTimestampNilField
			created_by: schema.AttributeUserRef
			updated_by: schema.AttributeUserRef
			deleted_by: schema.AttributeUserRef
		}

		indexes: {
			"primary": { attribute: "id" }
		}
	}

	envoy: {
		yaml: {
			supportMappedInput: false
			omitEncoder: true
		}
		store: {
			handleField: ""
		}
	}

	filter: {
		struct: {
			apigw_filter_id: {goType: "[]uint64", ident: "apigwFilterID", storeIdent: "id"}
			route_id: {goType: "uint64", ident: "routeID", storeIdent: "rel_route"}
			deleted:  {goType: "filter.State", storeIdent: "deleted_at"}
			disabled: {goType: "filter.State", storeIdent: "enabled"}
		}

		byValue: ["apigw_filter_id", "route_id"]
		byNilState: ["deleted"]
		byFalseState: ["disabled"]
	}

	store: {
		api: {
			lookups: [
				{
					fields: ["id"]
					description: """
						searches for filter by ID
						"""
				}, {
					fields: ["route"]
					description: """
						searches for filter by route
						"""
				},
			]
		}
	}
}
