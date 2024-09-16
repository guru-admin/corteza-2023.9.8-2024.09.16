package store

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/compose/types"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy/resource"
)

func newComposeNamespace(ns *types.Namespace) *composeNamespace {
	return &composeNamespace{
		ns: ns,
	}
}

func (ns *composeNamespace) MarshalEnvoy() ([]resource.Interface, error) {
	return envoy.CollectNodes(
		resource.NewComposeNamespace(ns.ns),
	)
}
