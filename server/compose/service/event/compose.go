package event

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/eventbus"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/scheduler"
)

// Match returns false if given conditions do not match event & resource internals
func (res composeOnInterval) Match(c eventbus.ConstraintMatcher) bool {
	return scheduler.OnInterval(c.Values()...)
}

// Match returns false if given conditions do not match event & resource internals
func (res composeOnTimestamp) Match(c eventbus.ConstraintMatcher) bool {
	return scheduler.OnTimestamp(c.Values()...)
}

// Match returns false if given conditions do not match event & resource internals
func (res composeBase) Match(c eventbus.ConstraintMatcher) bool {
	// We do not support any matchers, so if there is a constraint, fail.
	return false
}
