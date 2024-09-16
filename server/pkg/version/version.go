package version

import (
	"time"
)

var (
	// BuildTime value is set at build time and served over API and CLI
	BuildTime = time.Now().Format(time.RFC3339)

	// Version is set as LDFLAG at build time:
	// -X github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/version.Version=....
	// See Makefile for details
	Version = "development"
)
