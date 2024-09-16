package rdbms

import (
	"context"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/logger"
	"go.uber.org/zap"
)

// log() returns named logger with caller skip and stacktrace set to Fatal
//
// It checks the given context for logger before falling back to one set on the store
func (s Store) log(ctx context.Context) *zap.Logger {
	return logger.ContextValue(ctx, s.Logger, logger.Default(), zap.NewNop()).
		Named("store").
		WithOptions(zap.AddCallerSkip(2))
}

func (s Store) SetLogger(logger *zap.Logger) {
	s.Logger = logger
}
