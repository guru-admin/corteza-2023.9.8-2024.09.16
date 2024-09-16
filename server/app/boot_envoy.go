package app

import (
	"context"

	automationEnvoy "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/automation/envoy"
	composeEnvoy "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/compose/envoy"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoyx"
	systemEnvoy "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/envoy"
	"go.uber.org/zap"
)

func (app *CortezaApp) initEnvoy(ctx context.Context, log *zap.Logger) (err error) {
	// no-op - if envoy is already initialized
	if envoyx.Initialized() {
		return
	}

	// Init envoy
	svc := envoyx.New()
	envoyx.SetGlobal(svc)

	// Register all component decoders
	svc.AddDecoder(envoyx.DecodeTypeURI,
		composeEnvoy.YamlDecoder{},
		systemEnvoy.YamlDecoder{},
		automationEnvoy.YamlDecoder{},
	)
	svc.AddDecoder(envoyx.DecodeTypeStore,
		composeEnvoy.StoreDecoder{},
		systemEnvoy.StoreDecoder{},
		automationEnvoy.StoreDecoder{},
	)

	// Register all component encoders
	svc.AddEncoder(envoyx.EncodeTypeIo,
		composeEnvoy.YamlEncoder{},
		systemEnvoy.YamlEncoder{},
		automationEnvoy.YamlEncoder{},
	)
	svc.AddEncoder(envoyx.EncodeTypeStore,
		composeEnvoy.StoreEncoder{},
		systemEnvoy.StoreEncoder{},
		automationEnvoy.StoreEncoder{},
	)

	// - datasource encoders
	svc.AddEncoder(envoyx.EncodeTypeIo,
		composeEnvoy.CsvEncoder{},
	)

	return
}
