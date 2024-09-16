package envoy

import (
	"context"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoyx"
	"gopkg.in/yaml.v3"
)

func (e YamlEncoder) encode(ctx context.Context, base *yaml.Node, p envoyx.EncodeParams, rt string, nodes envoyx.NodeSet, tt envoyx.Traverser) (out *yaml.Node, err error) {
	return
}
