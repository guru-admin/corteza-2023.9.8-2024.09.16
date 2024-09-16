package store

import (
	"fmt"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy/resource"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/rbac"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
)

type (
	rbacRule struct {
		cfg *EncoderConfig

		rule *rbac.Rule

		// point to the rbac rule
		refRbacResource string
		refRbacRes      *resource.Ref

		refPathRes []*resource.Ref

		refRole *resource.Ref
		role    *types.Role
	}
)

func rbacResourceErrUnidentifiable(ii resource.Identifiers) error {
	return fmt.Errorf("rbac resource unidentifiable %v", ii.StringSlice())
}
