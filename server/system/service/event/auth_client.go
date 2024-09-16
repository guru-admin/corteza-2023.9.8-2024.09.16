package event

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/eventbus"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
)

var _ = eventbus.ConstraintMaker

// Match returns false if given conditions do not match event & resource internals
func (res authClientBase) Match(c eventbus.ConstraintMatcher) bool {
	return authClientMatch(res.authClient, c)
}

// Handles application matchers
func authClientMatch(r *types.AuthClient, c eventbus.ConstraintMatcher) bool {
	switch c.Name() {
	// not supported yet
	case "auth-client", "auth-client.handle":
		return c.Match(r.Handle)
	}

	return false
}
