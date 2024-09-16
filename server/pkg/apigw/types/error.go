package types

import (
	"net/http"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/errors"
	"go.uber.org/zap"
)

type (
	DefaultErrorHandler struct {
		log *zap.Logger
	}
)

func NewDefaultErrorHandler(log *zap.Logger) DefaultErrorHandler {
	return DefaultErrorHandler{
		log: log,
	}
}

func (h DefaultErrorHandler) Handler() ErrorHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request, err error) {
		errors.ProperlyServeHTTP(rw, r, err, true)
		h.log.Error(err.Error())
	}
}
