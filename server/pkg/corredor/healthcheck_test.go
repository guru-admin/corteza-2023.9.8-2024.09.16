package corredor

import (
	"context"
	"testing"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/options"
	"go.uber.org/zap"
)

// tested with
// go test -count 10 -race -run TestDataRace ./pkg/corredor/...
func TestDataRace(t *testing.T) {
	go Setup(zap.NewNop(), options.CorredorOpt{})
	go Healthcheck(context.Background())
}
