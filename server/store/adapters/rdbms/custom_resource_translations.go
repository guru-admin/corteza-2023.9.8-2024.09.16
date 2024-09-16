package rdbms

import (
	"context"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/locale"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
	"golang.org/x/text/language"
)

func (s Store) TransformResource(ctx context.Context, lang language.Tag) (out map[string]map[string]*locale.ResourceTranslation, err error) {
	out = make(map[string]map[string]*locale.ResourceTranslation)
	var cc types.ResourceTranslationSet
	var ok bool

	cc, _, err = s.SearchResourceTranslations(ctx, types.ResourceTranslationFilter{
		Lang: lang.String(),
	})
	if err != nil {
		return
	}

	for _, c := range cc {
		if _, ok = out[c.Resource]; !ok {
			out[c.Resource] = make(map[string]*locale.ResourceTranslation)
		}

		out[c.Resource][c.K] = &locale.ResourceTranslation{
			Resource: c.Resource,
			Lang:     c.Lang.String(),
			Key:      c.K,
			Msg:      c.Message,
		}
	}

	return
}
