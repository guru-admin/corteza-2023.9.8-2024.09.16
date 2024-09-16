package store

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy/resource"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
)

func newApplication(app *types.Application, ux *userIndex) *application {
	return &application{
		app: app,
		ux:  ux,
	}
}

func (app *application) MarshalEnvoy() ([]resource.Interface, error) {
	rs := resource.NewApplication(app.app)
	syncUserStamps(rs.Userstamps(), app.ux)

	return envoy.CollectNodes(
		rs,
	)
}
