package envoy

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoyx"
	"gopkg.in/yaml.v3"
)

func (d *auxYamlDoc) unmarshalTriggersExtendedNode(dctx documentContext, n *yaml.Node, meta ...*yaml.Node) (out envoyx.NodeSet, err error) {
	return d.unmarshalTriggerNode(dctx, n, meta...)
}

func (d *auxYamlDoc) unmarshalYAML(k string, n *yaml.Node) (out envoyx.NodeSet, err error) { return }
