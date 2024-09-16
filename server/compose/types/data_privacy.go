package types

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/filter"
	sysTypes "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
)

type (
	PrivacyModule struct {
		Module    PrivacyModuleMeta    `json:"module"`
		Namespace PrivacyNamespaceMeta `json:"namespace"`

		ConnectionID uint64                  `json:"-"`
		Connection   *sysTypes.DalConnection `json:"connection"`
	}

	PrivacyModuleMeta struct {
		ID     uint64         `json:"moduleID,string"`
		Name   string         `json:"name"`
		Handle string         `json:"handle"`
		Fields ModuleFieldSet `json:"fields"`
	}

	PrivacyNamespaceMeta struct {
		ID   uint64 `json:"namespaceID,string"`
		Slug string `json:"slug"`
		Name string `json:"name"`
	}

	PrivacyModuleFilter struct {
		NamespaceID  uint64   `json:"-"`
		ConnectionID []string `json:"connectionID"`

		// Standard helpers for paging and sorting
		filter.Sorting
		filter.Paging
	}
)
