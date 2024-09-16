package mssql

import (
	"context"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/dal"
	rdbmsdal "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/store/adapters/rdbms/dal"
)

func init() {
	// dal.RegisterConnector(dalConnector, SCHEMA, debugSchema)
	dal.RegisterConnector(dalConnector, SCHEMA)
}

func dalConnector(ctx context.Context, dsn string) (_ dal.Connection, err error) {
	cfg, err := NewConfig(dsn)
	if err != nil {
		return
	}

	// @todo rework the config building a bit; this will do for now
	if cfg.ConnTryMax >= 99 {
		cfg.ConnTryMax = 2
	}

	db, err := connectBase(ctx, cfg)

	if err != nil {
		return
	}
	return rdbmsdal.Connection(db, Dialect(), DataDefiner(cfg.DBName, db)), nil
}
