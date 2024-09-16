package yaml

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy/resource"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
)

type (
	role struct {
		res *types.Role
		ts  *resource.Timestamps

		envoyConfig   *resource.EnvoyConfig
		encoderConfig *EncoderConfig

		// module's RBAC rules
		rbac rbacRuleSet
	}
	roleSet []*role
)

func (nn roleSet) configureEncoder(cfg *EncoderConfig) {
	for _, n := range nn {
		n.encoderConfig = cfg
	}
}
