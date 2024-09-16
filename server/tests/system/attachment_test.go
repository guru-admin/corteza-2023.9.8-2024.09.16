package system

import (
	"context"
	"fmt"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"net/http"
	"testing"
	"time"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/id"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/store"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/service"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/tests/helpers"
)

func (h helper) clearAttachments() {
	h.noError(store.TruncateAttachments(context.Background(), service.DefaultStore))
}

func (h helper) repoMakeAttachment(ss ...string) *types.Attachment {
	var res = &types.Attachment{
		ID:        id.Next(),
		CreatedAt: time.Now(),
		Kind:      "json",
	}

	if len(ss) > 0 {
		res.Name = ss[0]
	} else {
		res.Name = "n_" + rs()
	}

	h.a.NoError(store.CreateAttachment(context.Background(), service.DefaultStore, res))

	return res
}

func (h helper) lookupAttachmentByID(ID uint64) *types.Attachment {
	res, err := store.LookupAttachmentByID(context.Background(), service.DefaultStore, ID)
	h.noError(err)
	return res
}

func TestAttachmentRead(t *testing.T) {
	h := newHelper(t)
	h.clearAttachments()

	a := h.repoMakeAttachment()

	h.apiInit().
		Get(fmt.Sprintf("/attachment/%s/%d", a.Kind, a.ID)).
		Expect(t).
		Status(http.StatusOK).
		Assert(helpers.AssertNoErrors).
		Assert(jsonpath.Equal(`$.response.name`, a.Name)).
		Assert(jsonpath.Equal(`$.response.attachmentID`, fmt.Sprintf("%d", a.ID))).
		End()
}

func TestAttachmentDelete(t *testing.T) {
	h := newHelper(t)
	h.clearAttachments()

	a := h.repoMakeAttachment()
	helpers.AllowMe(h, types.ApplicationRbacResource(0), "delete")

	h.apiInit().
		Delete(fmt.Sprintf("/attachment/%s/%d", a.Kind, a.ID)).
		Header("Accept", "application/json").
		Expect(t).
		Status(http.StatusOK).
		Assert(helpers.AssertNoErrors).
		End()

	a = h.lookupAttachmentByID(a.ID)
	h.a.NotNil(a)
	h.a.NotNil(a.DeletedAt)
}
