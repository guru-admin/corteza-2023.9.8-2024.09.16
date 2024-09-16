package postgres

import (
	"context"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/logger"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/store/adapters/rdbms"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/store/adapters/rdbms/drivers/postgres"
	"github.com/jmoiron/sqlx"
)

func Setup(ctx context.Context, dsn string) (_ *sqlx.DB, err error) {
	var (
		cfg *rdbms.ConnConfig
	)

	cfg, err = postgres.NewConfig(dsn)
	if err != nil {
		return
	}

	return rdbms.Connect(ctx, logger.MakeDebugLogger(), cfg)
}
