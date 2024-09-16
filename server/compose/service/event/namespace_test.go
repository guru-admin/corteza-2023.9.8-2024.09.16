package event

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/compose/types"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/eventbus"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNamespaceMatching(t *testing.T) {
	var (
		a   = assert.New(t)
		res = &namespaceBase{
			namespace: &types.Namespace{Slug: "slg1"},
		}

		cNms = eventbus.MustMakeConstraint("namespace", "eq", "slg1")
	)

	a.True(res.Match(cNms))
}
