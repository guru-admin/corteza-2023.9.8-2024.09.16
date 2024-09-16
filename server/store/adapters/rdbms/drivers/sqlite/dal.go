package sqlite

import (
	"context"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/dal"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/logger"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/store/adapters/rdbms"
	rdbmsdal "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/store/adapters/rdbms/dal"
	"github.com/jmoiron/sqlx"
)

func init() {
	dal.RegisterConnector(dalConnector, SCHEMA, altSchema, debugSchema)
}

func dalConnector(ctx context.Context, dsn string) (_ dal.Connection, err error) {
	var (
		db  *sqlx.DB
		cfg *rdbms.ConnConfig
	)

	if cfg, err = NewConfig(dsn); err != nil {
		return
	}

	// @todo rework the config building a bit; this will do for now
	if cfg.ConnTryMax >= 99 {
		cfg.ConnTryMax = 2
	}

	if db, err = rdbms.Connect(ctx, logger.Default(), cfg); err != nil {
		return
	}

	return rdbmsdal.Connection(db, Dialect(), DataDefiner(cfg.DBName, db)), nil
}
