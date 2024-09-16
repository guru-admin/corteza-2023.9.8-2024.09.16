package store

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy/resource"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
)

func newTemplate(t *types.Template) *template {
	return &template{
		t: t,
	}
}

func (u *template) MarshalEnvoy() ([]resource.Interface, error) {
	return envoy.CollectNodes(
		resource.NewTemplate(u.t),
	)
}
