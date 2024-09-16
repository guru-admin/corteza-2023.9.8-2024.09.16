package yaml

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy/resource"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
)

type (
	application struct {
		res *types.Application
		ts  *resource.Timestamps
		us  *resource.Userstamps

		envoyConfig   *resource.EnvoyConfig
		encoderConfig *EncoderConfig

		rbac rbacRuleSet
	}
	applicationSet []*application
)

func (nn applicationSet) configureEncoder(cfg *EncoderConfig) {
	for _, n := range nn {
		n.encoderConfig = cfg
	}
}
