package event

import "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/eventbus"

// Match returns false if given conditions do not match event & resource internals
func (res authBase) Match(c eventbus.ConstraintMatcher) bool {
	return userMatch(res.user, c)
}
