package event

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/eventbus"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoleMatching(t *testing.T) {
	var (
		a   = assert.New(t)
		res = &roleBase{
			role: &types.Role{Handle: "admin"},
		}

		cRol = eventbus.MustMakeConstraint("role", "eq", "admin")
	)

	a.True(res.Match(cRol))
}
