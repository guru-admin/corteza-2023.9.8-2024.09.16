package system

import (
	"net/http"
	"testing"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/tests/helpers"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func Test_dal_driver_list(t *testing.T) {
	h := newHelper(t)

	h.apiInit().
		Get("/dal/drivers/").
		Expect(t).
		Status(http.StatusOK).
		Assert(helpers.AssertNoErrors).
		Assert(jsonpath.Len("$.response.set", 1)).
		Assert(jsonpath.Present("$.response.set[0].operations")).
		End()
}
