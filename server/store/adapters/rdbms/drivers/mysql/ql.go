package mysql

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/store/adapters/rdbms/ql"
	"github.com/doug-martin/goqu/v9/exp"
)

var (
	ref2exp = ql.ExprHandlerMap{
		"std": {
			Handler: func(args ...exp.Expression) exp.Expression {
				return exp.NewSQLFunctionExpression("STD", args[0])
			},
		},
	}.ExprHandlers()
)
