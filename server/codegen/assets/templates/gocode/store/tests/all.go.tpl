package tests

{{ template "gocode/header-gentext.tpl" }}

import (
	"testing"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/store"
)

func testAllGenerated(t *testing.T, s store.Storer) {
{{ range .types }}
	t.Run({{ printf "%q" .ident }}, func(t *testing.T) {
		test{{ .expIdentPlural }}(t, s)
	})
{{- end }}
}
