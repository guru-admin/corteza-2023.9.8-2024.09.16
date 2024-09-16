package store

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy/resource"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
)

func newResourceTranslation(l types.ResourceTranslationSet) (*resourceTranslation, error) {
	res := l[0].Resource
	_, ref, pp, err := resource.ParseResourceTranslation(res)

	return &resourceTranslation{
		locales: l,

		refResourceTranslation: res,
		refLocaleRes:           ref,

		refPathRes: pp,
	}, err
}

func (lr *resourceTranslation) MarshalEnvoy() ([]resource.Interface, error) {
	return envoy.CollectNodes(
		resource.NewResourceTranslation(lr.locales, lr.refResourceTranslation, lr.refLocaleRes, lr.refPathRes...),
	)
}
