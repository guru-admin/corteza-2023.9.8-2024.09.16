package event

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/eventbus"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserMatching(t *testing.T) {
	var (
		a   = assert.New(t)
		res = &userBase{
			user: &types.User{Email: "user@example.tld"},
		}

		cUsr = eventbus.MustMakeConstraint("user.email", "eq", "user@example.tld")
	)

	a.True(res.Match(cUsr))
}
