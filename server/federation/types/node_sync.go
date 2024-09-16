package types

import (
	"time"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/filter"
)

var (
	NodeSyncTypeStructure = "sync_structure"
	NodeSyncTypeData      = "sync_data"
	NodeSyncStatusSuccess = "success"
	NodeSyncStatusError   = "error"
)

type (
	NodeSync struct {
		NodeID     uint64 `json:"nodeID,string"`
		ModuleID   uint64 `json:"moduleID,string"`
		SyncStatus string `json:"syncStatus"`
		SyncType   string `json:"syncType"`

		TimeOfAction time.Time `json:"timeOfAction"`
	}

	NodeSyncFilter struct {
		NodeID     uint64 `json:"nodeID"`
		ModuleID   uint64 `json:"moduleID"`
		SyncStatus string `json:"syncStatus"`
		SyncType   string `json:"syncType"`

		Query string `json:"name"`

		Check func(*NodeSync) (bool, error) `json:"-"`

		filter.Sorting
		filter.Paging
	}
)
