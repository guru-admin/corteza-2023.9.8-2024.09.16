package yaml

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy/resource"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
)

type (
	setting struct {
		res *types.SettingValue
		ts  *resource.Timestamps
		us  *resource.Userstamps

		envoyConfig   *resource.EnvoyConfig
		encoderConfig *EncoderConfig
	}

	settingSet []*setting
)
