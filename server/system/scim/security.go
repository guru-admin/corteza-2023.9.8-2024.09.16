package scim

import (
	"context"
	"net/http"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/auth"
)

type (
	getSecurityContextFn func(r *http.Request) context.Context
)

// Set service user to request's identity
func getSecurityContext(r *http.Request) context.Context {
	return auth.SetIdentityToContext(r.Context(), auth.ServiceUser())
}
