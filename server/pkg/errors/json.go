package errors

import (
	"encoding/json"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/locale"
)

func (e Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Message string   `json:"message"`
		Meta    meta     `json:"meta,omitempty"`
		Stack   []*frame `json:"stack,omitempty"`
		Wrap    error    `json:"wrap,omitempty"`
	}{
		Message: e.Error(),
		Meta:    e.meta,
		Stack:   e.stack,
		Wrap:    e.wrap,
	})
}

func (e Error) Translate(tr func(string, string, ...string) string) error {
	if msg := tr(e.meta.AsString(locale.ErrorMetaNamespace{}), e.meta.AsString(locale.ErrorMetaKey{}), e.meta.pairs()...); len(msg) > 0 {
		e.message = msg
	}

	return &e
}
