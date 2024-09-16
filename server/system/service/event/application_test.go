package event

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/eventbus"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplicationMatching(t *testing.T) {
	var (
		a   = assert.New(t)
		res = &applicationBase{
			application: &types.Application{Name: "someApp"},
		}

		cApp = eventbus.MustMakeConstraint("application.name", "eq", "someApp")
	)

	a.True(res.Match(cApp))
}
