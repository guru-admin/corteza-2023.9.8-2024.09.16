package service

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/cli"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/id"
)

func init() {
	id.Init(cli.Context())
}
