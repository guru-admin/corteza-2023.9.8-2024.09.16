package store

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy/resource"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
)

func newSetting(res *types.SettingValue, ux *userIndex) *setting {
	return &setting{
		st: res,

		ux: ux,
	}
}

func (st *setting) MarshalEnvoy() ([]resource.Interface, error) {
	rs := resource.NewSetting(st.st)
	syncUserStamps(rs.Userstamps(), st.ux)

	return envoy.CollectNodes(
		rs,
	)
}
