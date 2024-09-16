package system

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/actionlog"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/id"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/store"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/service"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/tests/helpers"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func (h helper) clearActionLog() {
	h.noError(store.TruncateActionlogs(context.Background(), service.DefaultStore))
}

func (h helper) repoMakeActionLog() *actionlog.Action {
	var res = &actionlog.Action{
		ID:        id.Next(),
		Timestamp: time.Now(),
		ActorID:   id.Next(),
		Resource:  types.ComponentRbacResource(),
		Action:    "lookup",
	}

	h.a.NoError(store.CreateActionlog(context.Background(), service.DefaultStore, res))

	return res
}

func TestActionLogList(t *testing.T) {
	h := newHelper(t)
	h.clearActionLog()

	helpers.AllowMe(h, types.ComponentRbacResource(), "action-log.read")

	h.repoMakeActionLog()

	h.apiInit().
		Get("/actionlog/").
		Header("Accept", "application/json").
		Expect(t).
		Status(http.StatusOK).
		Assert(helpers.AssertNoErrors).
		Assert(jsonpath.Len(`$.response.set`, 1)).
		End()
}

func TestActionLogListForbidden(t *testing.T) {
	h := newHelper(t)
	h.clearActionLog()

	h.repoMakeActionLog()

	h.apiInit().
		Get("/actionlog/").
		Header("Accept", "application/json").
		Expect(t).
		Status(http.StatusOK).
		Assert(helpers.AssertError("access_denied")).
		End()
}
