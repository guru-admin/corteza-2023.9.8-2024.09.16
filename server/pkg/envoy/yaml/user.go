package yaml

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy/resource"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
)

type (
	user struct {
		res   *types.User
		ts    *resource.Timestamps
		roles []string

		envoyConfig   *resource.EnvoyConfig
		encoderConfig *EncoderConfig

		rbac rbacRuleSet
	}
	userSet []*user
)

func (nn userSet) configureEncoder(cfg *EncoderConfig) {
	for _, n := range nn {
		n.encoderConfig = cfg
	}
}
