package {{ .package }}

{{ template "gocode/header-gentext.tpl" }}

import (
	"context"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/dal"
{{- range .imports }}
    {{ . }}
{{- end }}
)


type (
	modelReplacer interface  {
		ReplaceModel(ctx context.Context, model *dal.Model) (err error)
	}
)

var (
	models []*dal.Model
)

func Models() dal.ModelSet {
	return models
}
