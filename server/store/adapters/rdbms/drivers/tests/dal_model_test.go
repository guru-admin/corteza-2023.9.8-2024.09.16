package tests

import (
	"context"
	"testing"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/dal"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/logger"
	"github.com/stretchr/testify/require"
)

func TestModelManagement(t *testing.T) {
	var (
		ctx = context.Background()
		log = logger.Default()
	)

	var (
		req = require.New(t)
	)

	if conn.isSQLite {
		t.Skip("sqlite does not support index management")
	}

	const (
		dalConnID    = 1
		dalTableName = "lil_dal_test"
	)

	_, err := conn.db.Exec("DROP TABLE IF EXISTS " + dalTableName)
	req.NoError(err)

	cw := dal.MakeConnection(
		dalConnID,
		conn.store.ToDalConn(),
		dal.ConnectionParams{},
		dal.ConnectionConfig{},
	)

	svc, err := dal.New(log, true)
	req.NoError(err)
	req.NoError(svc.ReplaceConnection(ctx, cw, true))

	_, err = svc.ReplaceModel(ctx, nil, &dal.Model{
		Ident:        dalTableName,
		ConnectionID: dalConnID,
		Attributes: dal.AttributeSet{
			&dal.Attribute{
				Ident:      "ID",
				PrimaryKey: true,
				Type:       &dal.TypeID{},
				Store:      &dal.CodecAlias{Ident: "id"},
			},

			&dal.Attribute{
				Ident:    "OwnerID",
				Sortable: true,
				Type:     &dal.TypeText{},
				Store:    &dal.CodecAlias{Ident: "rel_owner"},
			},
		},
	})
	req.NoError(err)

	_, err = conn.db.Exec("DROP TABLE " + dalTableName)
	req.NoError(err)
}
