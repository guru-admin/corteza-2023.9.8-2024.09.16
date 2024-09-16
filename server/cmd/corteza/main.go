package main

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/app"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/cli"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/logger"
)

func main() {
	// Initialize logger before any other action
	logger.Init()

	cli.HandleError(app.New().Execute())
}
