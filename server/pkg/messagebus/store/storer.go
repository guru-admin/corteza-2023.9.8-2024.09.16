package store

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/messagebus/types"
)

type (
	Storer interface {
		SetStore(types.QueueStorer)
		GetStore() types.QueueStorer
	}
)
