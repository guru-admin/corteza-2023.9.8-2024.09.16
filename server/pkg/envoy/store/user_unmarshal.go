package store

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy/resource"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
)

func newUser(u *types.User) *user {
	return &user{
		u: u,
	}
}

func (u *user) MarshalEnvoy() ([]resource.Interface, error) {
	return envoy.CollectNodes(
		resource.NewUser(u.u),
	)
}
