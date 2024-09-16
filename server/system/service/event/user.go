package event

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/eventbus"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
)

// Match returns false if given conditions do not match event & resource internals
func (res userBase) Match(c eventbus.ConstraintMatcher) bool {
	return userMatch(res.user, c)
}

// Handles user matchers
func userMatch(r *types.User, c eventbus.ConstraintMatcher) bool {
	switch c.Name() {
	case "user", "user.handle":
		return r != nil && c.Match(r.Handle)
	case "user.email":
		return r != nil && c.Match(r.Email)
	}

	return false
}
