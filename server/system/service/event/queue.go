package event

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/eventbus"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
)

var _ = eventbus.ConstraintMaker

// Match returns false if given conditions do not match event & resource internals
func (res queueBase) Match(c eventbus.ConstraintMatcher) bool {
	return queueMatch(res.payload, c)
}

func queueMatch(r *types.QueueMessage, c eventbus.ConstraintMatcher) bool {
	switch c.Name() {
	case "payload.queue":
		return c.Match(r.Queue)
	}

	return false
}
