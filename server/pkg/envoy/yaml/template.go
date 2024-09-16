package yaml

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy/resource"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
)

type (
	template struct {
		res *types.Template
		ts  *resource.Timestamps
		us  *resource.Userstamps

		envoyConfig   *resource.EnvoyConfig
		encoderConfig *EncoderConfig

		rbac rbacRuleSet
	}
	templateSet []*template
)

func (nn templateSet) configureEncoder(cfg *EncoderConfig) {
	for _, n := range nn {
		n.encoderConfig = cfg
	}
}
