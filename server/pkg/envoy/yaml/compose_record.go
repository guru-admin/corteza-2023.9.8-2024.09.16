package yaml

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/compose/types"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy/resource"
)

type (
	composeRecord struct {
		id     string
		values map[string]string
		ts     *resource.Timestamps
		us     *resource.Userstamps
		config *resource.EnvoyConfig

		cfg *EncoderConfig

		refModule    string
		refNamespace string

		rbac rbacRuleSet
	}
	composeRecordSet []*composeRecord

	composeRecordValues struct {
		rvs types.RecordValueSet
	}
)

func (nn composeRecordSet) configureEncoder(cfg *EncoderConfig) {
	for _, n := range nn {
		n.cfg = cfg
	}
}
