package event

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/eventbus"
)

// Match returns false if given conditions do not match event & resource internals
func (res roleMemberBase) Match(c eventbus.ConstraintMatcher) bool {
	return eventbus.MatchFirst(
		func() bool { return userMatch(res.user, c) },
		func() bool { return roleMatch(res.role, c) },
	)
}
