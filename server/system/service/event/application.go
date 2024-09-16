package event

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/eventbus"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
)

// Match returns false if given conditions do not match event & resource internals
func (res applicationBase) Match(c eventbus.ConstraintMatcher) bool {
	return applicationMatch(res.application, c)
}

// Handles application matchers
func applicationMatch(r *types.Application, c eventbus.ConstraintMatcher) bool {
	switch c.Name() {
	// not supported yet
	//case "application", "application.handle":
	//	return c.Match(r.Handle)
	case "application.name":
		return c.Match(r.Name)

	}

	return false
}
