package types

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/filter"
)

type (
	ModuleMapping struct {
		NodeID             uint64                `json:"nodeID,string"`
		FederationModuleID uint64                `json:"federationModuleID,string"`
		ComposeModuleID    uint64                `json:"composeModuleID,string"`
		ComposeNamespaceID uint64                `json:"composeNamespaceID,string"`
		FieldMapping       ModuleFieldMappingSet `json:"fields"`
	}

	ModuleMappingFilter struct {
		NodeID             uint64 `json:"nodeID"`
		ComposeModuleID    uint64 `json:"composeModuleID"`
		ComposeNamespaceID uint64 `json:"composeNamespaceID"`
		FederationModuleID uint64 `json:"federationModuleID"`
		Query              string `json:"query"`

		Check func(*ModuleMapping) (bool, error) `json:"-"`

		filter.Sorting
		filter.Paging
	}
)
