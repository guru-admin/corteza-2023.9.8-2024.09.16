package event

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/eventbus"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoleMemberMatching(t *testing.T) {
	var (
		a   = assert.New(t)
		res = &roleMemberBase{
			role: &types.Role{Handle: "admin"},
			user: &types.User{Handle: "user"},
		}

		cRol = eventbus.MustMakeConstraint("role", "eq", "admin")
		cUsr = eventbus.MustMakeConstraint("user", "eq", "user")
	)

	a.True(res.Match(cRol))
	a.True(res.Match(cUsr))
}
