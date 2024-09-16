package store

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy/resource"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
)

type (
	setting struct {
		cfg *EncoderConfig

		res *resource.Setting
		st  *types.SettingValue

		ux *userIndex
	}
)
