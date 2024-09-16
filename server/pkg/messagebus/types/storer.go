package types

import (
	"context"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/eventbus"
)

type (
	QueueStorer interface {
		SearchQueues(context.Context, QueueFilter) ([]QueueDb, QueueFilter, error)
		CreateQueueMessage(context.Context, QueueMessage) error
		ProcessQueueMessage(context.Context, uint64, QueueMessage) error
	}

	QueueEventBuilder interface {
		CreateQueueEvent(string, []byte) eventbus.Event
	}

	QueueServicer interface {
		QueueStorer
		QueueEventBuilder
	}
)
