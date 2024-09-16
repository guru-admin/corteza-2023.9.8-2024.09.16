package rest

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/auth"
	"net/http"
)

type Sink struct {
	svc interface {
		ProcessRequest(w http.ResponseWriter, r *http.Request)
	}

	sign auth.Signer
}

func (ctrl *Sink) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctrl.svc.ProcessRequest(w, r)
}
