package decoder

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/compose/types"
)

type (
	ComposeRecord struct {
		types.Record
	}
	ComposeRecordSet []*ComposeRecord

	ComposeRecordFilter struct {
		types.RecordFilter
	}
)
