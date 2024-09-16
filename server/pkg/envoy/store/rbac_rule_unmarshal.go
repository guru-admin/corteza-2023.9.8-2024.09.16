package store

import (
	"strconv"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy/resource"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/rbac"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
)

func newRbacRule(rl *rbac.Rule) (*rbacRule, error) {
	res := rl.Resource
	_, ref, pp, err := resource.ParseRule(res)

	return &rbacRule{
		rule: rl,

		refRbacResource: res,
		refRbacRes:      ref,

		refPathRes: pp,

		refRole: resource.MakeRef(types.RoleResourceType, resource.MakeIdentifiers(strconv.FormatUint(rl.RoleID, 10))),
	}, err
}

func (rl *rbacRule) MarshalEnvoy() ([]resource.Interface, error) {
	return envoy.CollectNodes(
		resource.NewRbacRule(rl.rule, rl.refRole, rl.refRbacRes, rl.refRbacResource, rl.refPathRes...),
	)
}
