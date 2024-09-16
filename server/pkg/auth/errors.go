package auth

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/errors"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/locale"
)

func errUnauthorized() error {
	return errors.New(
		errors.KindUnauthorized,

		"unauthorized",

		errors.Meta("type", "unauthorized"),

		// translation namespace & key
		errors.Meta(locale.ErrorMetaNamespace{}, "internal"),
		errors.Meta(locale.ErrorMetaKey{}, "auth.errors.unauthorized"),

		errors.StackSkip(1),
		errors.StackTrimAtFn("http.HandlerFunc.ServeHTTP"),
	)
}

func errUnauthorizedScope() error {
	return errors.New(
		errors.KindUnauthorized,

		"unauthorized scope",

		errors.Meta("type", "unauthorizedScope"),

		// translation namespace & key
		errors.Meta(locale.ErrorMetaNamespace{}, "internal"),
		errors.Meta(locale.ErrorMetaKey{}, "auth.errors.unauthorizedScope"),

		errors.StackSkip(1),
		errors.StackTrimAtFn("http.HandlerFunc.ServeHTTP"),
	)
}
