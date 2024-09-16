package store

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/compose/types"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy/resource"
)

func newComposeModule(mod *types.Module) *composeModule {
	return &composeModule{
		mod: mod,
	}
}

func (mod *composeModule) MarshalEnvoy() ([]resource.Interface, error) {
	refNs := resource.MakeNamespaceRef(mod.mod.NamespaceID, "", "")
	refMod := resource.MakeModuleRef(mod.mod.ID, mod.mod.Handle, mod.mod.Name)

	rMod := resource.NewComposeModule(mod.mod, refNs)
	for _, f := range mod.mod.Fields {
		r := resource.NewComposeModuleField(f, refNs, refMod)
		rMod.AddField(r)
	}

	return envoy.CollectNodes(
		rMod,
	)
}
