package service

import (
	"context"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/eventbus"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/options"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/store"
	"go.uber.org/zap"
)

var (
	DefaultOption options.DiscoveryOpt

	// DefaultStore is an interface to storage backend(s)
	// ng (next-gen) is a temporary prefix
	// so that we can differentiate between it and the file-only store
	DefaultStore store.Storer

	DefaultResourceActivity *resourceActivity
)

// Initialize discovery service
func Initialize(_ context.Context, log *zap.Logger, opt options.DiscoveryOpt, s store.Storer) (err error) {
	DefaultOption = opt

	// we're doing conversion to avoid having
	// store interface exposed or generated inside app package
	DefaultStore = s

	DefaultResourceActivity = ResourceActivity(log, opt, s, eventbus.Service())

	return
}
