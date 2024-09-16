package rest

import (
	"context"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/automation/rest/request"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/automation/service"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/automation/types"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/api"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/rbac"
)

type (
	Permissions struct {
		ac permissionsAccessController
	}

	permissionsAccessController interface {
		Effective(context.Context, ...rbac.Resource) rbac.EffectiveSet
		Trace(context.Context, uint64, []uint64, ...string) ([]*rbac.Trace, error)
		List() []map[string]string
		FindRulesByRoleID(context.Context, uint64) (rbac.RuleSet, error)
		FindRules(ctx context.Context, roleID uint64, rr ...string) (rbac.RuleSet, error)
		Grant(ctx context.Context, rr ...*rbac.Rule) error
	}
)

func (Permissions) New() *Permissions {
	return &Permissions{
		ac: service.DefaultAccessControl,
	}
}

func (ctrl Permissions) Effective(ctx context.Context, r *request.PermissionsEffective) (interface{}, error) {
	return ctrl.ac.Effective(ctx, types.Component{}), nil
}

func (ctrl Permissions) Trace(ctx context.Context, r *request.PermissionsTrace) (interface{}, error) {
	return ctrl.ac.Trace(ctx, r.UserID, r.RoleID, r.Resource...)
}

func (ctrl Permissions) List(ctx context.Context, r *request.PermissionsList) (interface{}, error) {
	return ctrl.ac.List(), nil
}

func (ctrl Permissions) Read(ctx context.Context, r *request.PermissionsRead) (interface{}, error) {
	return ctrl.ac.FindRules(ctx, r.RoleID, r.Resource...)
}

func (ctrl Permissions) Delete(ctx context.Context, r *request.PermissionsDelete) (interface{}, error) {
	rr, err := ctrl.ac.FindRulesByRoleID(ctx, r.RoleID)
	if err != nil {
		return nil, err
	}

	for _, r := range rr {
		r.Access = rbac.Inherit
	}

	return api.OK(), ctrl.ac.Grant(ctx, rr...)
}

func (ctrl Permissions) Update(ctx context.Context, r *request.PermissionsUpdate) (interface{}, error) {
	for _, rule := range r.Rules {
		// Make sure everything is properly set
		rule.RoleID = r.RoleID
	}

	return api.OK(), ctrl.ac.Grant(ctx, r.Rules...)
}
