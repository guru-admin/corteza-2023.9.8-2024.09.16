package event

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/compose/types"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/eventbus"
)

// Match returns false if given conditions do not match event & resource internals
func (res namespaceBase) Match(c eventbus.ConstraintMatcher) bool {
	return namespaceMatch(res.namespace, c)
}

// Handles namespace matchers
func namespaceMatch(r *types.Namespace, c eventbus.ConstraintMatcher) bool {
	switch c.Name() {
	case "namespace", "namespace.slug", "namespace.handle":
		return c.Match(r.Slug)
	case "namespace.name":
		return c.Match(r.Name)
	}

	return false
}
