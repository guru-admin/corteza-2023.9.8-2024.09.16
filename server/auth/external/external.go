package external

import (
	"github.com/gorilla/sessions"
	"github.com/markbates/goth/gothic"
)

const (
	OIDC_PROVIDER_PREFIX = "openid-connect." // must match const in "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types" app_settings.go

)

func Init(store sessions.Store) {
	gothic.Store = store
}
