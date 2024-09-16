package system

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/codegen/schema"
)

queue_message: {
	features: {
		labels: false
		checkFn: false
	}

	model: {
		omitGetterSetter: true

		attributes: {
		  id:        schema.IdField
		  queue:     {
		  	sortable: true
		  	dal: {}
		  }
		  payload:   {
		  	goType: "[]byte"
		  	dal: { type: "Blob" }
		  }
		  created:   schema.SortableTimestampNilField
		  processed: schema.SortableTimestampNilField
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
			queue: {}
			processed: {goType: "filter.State", storeIdent: "processed"}
		}

		byValue: ["queue"]
		byNilState: ["processed"]
	}

	store: {
		api: {
			lookups: []
		}
	}
}
