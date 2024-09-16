package app

// Registers all supported store backends
import (
	_ "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/store/adapters/rdbms/drivers/mssql"
	_ "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/store/adapters/rdbms/drivers/mysql"
	_ "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/store/adapters/rdbms/drivers/postgres"
	_ "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/store/adapters/rdbms/drivers/sqlite"
)
