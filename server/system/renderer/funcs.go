package renderer

import "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/valuestore"

func envGetter() func(k string) any {
	return func(k string) any {
		return valuestore.Global().Env(k)
	}
}
