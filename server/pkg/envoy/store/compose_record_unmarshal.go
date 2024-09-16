package store

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy/resource"
)

func newComposeRecordFromAux(rec *composeRecordAux) *composeRecord {
	return &composeRecord{
		rec:    rec,
		relMod: rec.relMod,
	}
}

func (rec *composeRecord) MarshalEnvoy() ([]resource.Interface, error) {
	rr := resource.NewComposeRecordSet(rec.rec.walker, rec.rec.refNs, rec.rec.refMod)
	rr.SetUserFlakes(rec.rec.relUsers)
	rr.RelMod = rec.relMod

	return envoy.CollectNodes(
		rr,
	)
}
