package system

import (
	"context"
	discoveryType "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/discovery/types"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/id"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/store"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/service"
	"testing"
)

func (h helper) clearActivityLog() {
	h.noError(store.TruncateResourceActivitys(context.Background(), service.DefaultStore))
}

func (h helper) repoMakeActivityLog() *discoveryType.ResourceActivity {
	var res = &discoveryType.ResourceActivity{
		ID:             id.Next(),
		ResourceID:     id.Next(),
		ResourceType:   "compose:record",
		ResourceAction: "create",
	}

	h.a.NoError(store.CreateResourceActivity(context.Background(), service.DefaultStore, res))

	return res
}

func TestCreateActivityLog(t *testing.T) {
	h := newHelper(t)
	h.clearActionLog()

	h.repoMakeActivityLog()
}
