package rest

import (
	"context"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/federation/rest/request"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/federation/service"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/api"
)

type (
	handshakeInitializer interface {
		HandshakeInit(context.Context, uint64, string, uint64, string) error
	}

	NodeHandshake struct {
		svcNode handshakeInitializer
	}
)

func (NodeHandshake) New() *NodeHandshake {
	return &NodeHandshake{
		svcNode: service.DefaultNode,
	}
}

func (ctrl NodeHandshake) Initialize(ctx context.Context, r *request.NodeHandshakeInitialize) (interface{}, error) {
	return api.OK(), ctrl.svcNode.HandshakeInit(ctx, r.NodeID, r.PairToken, r.SharedNodeID, r.AuthToken)
}
