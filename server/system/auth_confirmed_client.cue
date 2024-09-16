package system

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/codegen/schema"
)

auth_confirmed_client: {
	features: {
		labels: false
		paging: false
		sorting: false
		checkFn: false
	}

	model: {
		omitGetterSetter: true

		attributes: {
			user_id: {
				goType: "uint64",
				ident: "userID",
				storeIdent: "rel_user"
				dal: { type: "Ref", refModelResType: "corteza::system:user", default: 0 }
		  }
			client_id: {
				goType: "uint64",
				ident: "clientID",
				storeIdent: "rel_client"
				dal: { type: "Ref", refModelResType: "corteza::system:auth-client", default: 0 }
			}
			confirmed_at: schema.SortableTimestampNowField
		}

		indexes: {
			"primary": { attributes: ["user_id", "client_id"] }
		}
	}

	filter: {
		struct: {
			user_id:   { goType: "uint64", ident: "userID", storeIdent: "rel_user" }
		}

		byValue: ["user_id"]
	}

	envoy: {
		omit: true
	}


	store: {
		api: {
			lookups: [
				{ fields: ["user_id", "client_id"] },
			]
		}
	}
}
