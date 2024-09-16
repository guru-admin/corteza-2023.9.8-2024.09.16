package external

import (
	"os"
	"testing"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/logger"
)

func TestMain(m *testing.M) {
	logger.SetDefault(logger.MakeDebugLogger())
	os.Exit(m.Run())
}
