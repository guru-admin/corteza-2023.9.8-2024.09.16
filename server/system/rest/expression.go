package rest

import (
	"context"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/rest/request"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/service"
)

type (
	Expression struct {
		svc exprService
	}

	exprService interface {
		Evaluate(context.Context, map[string]string, map[string]any) (map[string]any, error)
	}
)

func (Expression) New() *Expression {
	return &Expression{
		svc: service.DefaultExpression,
	}
}

func (ctrl *Expression) Evaluate(ctx context.Context, r *request.ExpressionEvaluate) (interface{}, error) {
	return ctrl.svc.Evaluate(ctx, r.Expressions, r.Variables)
}
