package automation

import (
	"context"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/expr"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUser(t *testing.T) {
	var (
		req    = require.New(t)
		u, err = NewUser(&types.User{Handle: "handle"})
	)

	req.NoError(err)
	req.Equal("handle", u.value.Handle)
	req.Error(u.AssignFieldValue("some-unexisting-field", nil))
	req.NoError(u.AssignFieldValue("email", expr.Must(expr.NewString("dummy@domain.tpl"))))
	req.Equal("dummy@domain.tpl", u.value.Email)
}

func TestUser_Expr(t *testing.T) {
	var (
		req   = require.New(t)
		u, _  = NewUser(&types.User{Handle: "hendl"})
		scope = &expr.Vars{}
	)

	req.NoError(scope.Set("user", u))

	eval, err := expr.NewParser().Parse("user.handle")
	req.NoError(err)

	res, err := eval.Eval(context.Background(), scope)
	req.NoError(err)

	req.Equal("hendl", res.(string))
}
