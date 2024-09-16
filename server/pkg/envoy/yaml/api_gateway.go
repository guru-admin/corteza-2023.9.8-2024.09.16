package yaml

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy/resource"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
)

type (
	apiGateway struct {
		res     *types.ApigwRoute
		filters apiGwFilterSet

		ts *resource.Timestamps
		us *resource.Userstamps

		envoyConfig   *resource.EnvoyConfig
		encoderConfig *EncoderConfig

		rbac rbacRuleSet
	}
	apiGatewaySet []*apiGateway

	apiGwFilter struct {
		res *types.ApigwFilter

		ts *resource.Timestamps
		us *resource.Userstamps

		envoyConfig   *resource.EnvoyConfig
		encoderConfig *EncoderConfig
	}
	apiGwFilterSet []*apiGwFilter
)

func (nn apiGatewaySet) configureEncoder(cfg *EncoderConfig) {
	for _, n := range nn {
		n.encoderConfig = cfg
	}
}
