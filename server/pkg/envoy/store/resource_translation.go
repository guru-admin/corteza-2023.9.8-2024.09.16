package store

import (
	"fmt"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy/resource"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
)

type (
	resourceTranslation struct {
		cfg *EncoderConfig

		locales types.ResourceTranslationSet

		// point to the resource translation
		refResourceTranslation string
		refLocaleRes           *resource.Ref

		refPathRes []*resource.Ref
	}
)

func resourceTranslationErrUnidentifiable(ii resource.Identifiers) error {
	return fmt.Errorf("resource translation unidentifiable %v", ii.StringSlice())
}
