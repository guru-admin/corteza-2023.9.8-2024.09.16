package event

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/eventbus"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthMatching(t *testing.T) {
	var (
		a   = assert.New(t)
		res = &authBase{
			user: &types.User{Handle: "user"},
		}

		cUsr = eventbus.MustMakeConstraint("user", "eq", "user")
	)

	a.True(res.Match(cUsr))
}
