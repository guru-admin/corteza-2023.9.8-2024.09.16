package corredor

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
)

type (
	userLookupCache struct {
		err  error
		user *types.User
	}

	userLookupCacheMap map[string]userLookupCache
)

func (m userLookupCacheMap) lookup(key string, lookup func() (*types.User, error)) (*types.User, error) {
	if c, ok := m[key]; ok {
		return c.user, c.err
	}

	c := userLookupCache{}
	c.user, c.err = lookup()

	m[key] = c

	return c.user, c.err
}
