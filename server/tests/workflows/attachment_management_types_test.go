package workflows

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/automation/types"
	cmpTypes "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/compose/types"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/expr"
	"github.com/stretchr/testify/require"
)

type (
	auxReader struct {
		read bool
	}
)

func Test_attachment_management_types(t *testing.T) {
	var (
		ctx = bypassRBAC(context.Background())
		req = require.New(t)
	)

	req.NoError(defStore.TruncateAttachments(ctx))

	loadNewScenario(ctx, t)

	var (
		aux = struct {
			AttachedString     *cmpTypes.Attachment
			AttachedReader     *cmpTypes.Attachment
			AttachedReadSeeker *cmpTypes.Attachment
			AttachedBytes      *cmpTypes.Attachment
		}{}
	)

	v := &expr.Vars{}
	v.AssignFieldValue("sourceString", expr.Must(expr.NewString("hello")))
	v.AssignFieldValue("sourceReader", expr.Must(expr.NewReader(&auxReader{})))
	v.AssignFieldValue("sourceReadSeeker", expr.Must(expr.NewReader(strings.NewReader("hello"))))
	v.AssignFieldValue("sourceBytes", expr.Must(expr.NewBytes([]byte{'h', 'e', 'l', 'l', 'o'})))

	vars, _ := mustExecWorkflow(ctx, t, "attachments", types.WorkflowExecParams{
		Input: v,
	})
	req.NoError(vars.Decode(&aux))

	req.NotNil(aux.AttachedString)
	req.NotNil(aux.AttachedReader)
	req.NotNil(aux.AttachedReadSeeker)
	req.NotNil(aux.AttachedBytes)
}

func (ar *auxReader) Read(dst []byte) (int, error) {
	if ar.read {
		return 0, io.EOF
	}

	for i := range dst {
		dst[i] = byte('a')
	}

	ar.read = true
	return len(dst), nil
}
