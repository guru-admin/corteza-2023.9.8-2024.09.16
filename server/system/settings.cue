package system

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/codegen/schema"
)

settings: {
	ident: "settingValue"
	expIdent: "SettingValue"

	features: {
		labels: false
		paging: false
		sorting: false
		checkFn: false
	}

	model: {
		ident: "settings"
		omitGetterSetter: true
		attributes: {
			owned_by:    {
				goType: "uint64",
				storeIdent: "rel_owner"
				dal: { type: "Ref", refModelResType: "corteza::system:user" }
			}
			name:        {
				dal: { type: "Text", length: 512 }
			}
			value:       {
				goType: "rawJson"
				dal: { type: "JSON" }
				omitSetter: true
				omitGetter: true
			}
			updated_by:  schema.AttributeUserRef
			updated_at:  schema.SortableTimestampField
		}

		indexes: {
			"primary": {
				fields: [
					{ attribute: "owned_by" },
				 	{ attribute: "name" },
				]
			}
		}
	}

	filter: {
		expIdent: "SettingsFilter"

		struct: {
			prefix: {}
			owned_by: { goType: "uint64", storeIdent: "rel_owner" }
		}

		byValue: [ "owned_by" ]
	}

	envoy: {
		omit: true
	}

	store: {
		api: {
			lookups: [
				{
					fields: ["name", "owned_by"]
					description: """
						searches for settings by name and owner
						"""
				}
			]
		}
	}
}
