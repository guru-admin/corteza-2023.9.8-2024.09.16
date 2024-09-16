package consumer

import (
	"context"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/eventbus"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/messagebus/types"
)

type (
	Dispatcher interface {
		Dispatch(ctx context.Context, ev eventbus.Event)
		WaitFor(ctx context.Context, ev eventbus.Event) (err error)
	}

	EventbusConsumer struct {
		queue      string
		handle     types.ConsumerType
		dispatcher Dispatcher
		servicer   types.QueueEventBuilder
	}
)

func NewEventbusConsumer(q string, servicer types.QueueEventBuilder) *EventbusConsumer {
	h := &EventbusConsumer{
		queue:      q,
		handle:     types.ConsumerEventbus,
		dispatcher: eventbus.Service(),
		servicer:   servicer,
	}

	return h
}

func (cq *EventbusConsumer) Write(ctx context.Context, p []byte) error {
	cq.dispatcher.Dispatch(ctx, cq.servicer.CreateQueueEvent(cq.queue, p))
	return nil
}
