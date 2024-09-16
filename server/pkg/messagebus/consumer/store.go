package consumer

import (
	"context"
	"time"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/messagebus/store"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/messagebus/types"
)

type (
	StoreConsumer struct {
		queue  string
		handle types.ConsumerType
		client store.StoreClient
		poll   *time.Ticker
	}
)

func NewStoreConsumer(q string, s types.QueueServicer) *StoreConsumer {
	h := &StoreConsumer{
		queue:  q,
		handle: types.ConsumerStore,
		client: store.NewClient(s),
	}

	return h
}

func (cq *StoreConsumer) Write(ctx context.Context, p []byte) error {
	return cq.client.Add(ctx, cq.queue, p)
}

func (cq *StoreConsumer) SetStore(s types.QueueStorer) {
	cq.client.SetStore(s)
}

func (cq *StoreConsumer) GetStore() types.QueueStorer {
	return cq.client.GetStore()
}
