package rest

import (
	"context"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/discovery/rest/internal/feed"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/discovery/rest/request"
	"time"
)

type (
	feeder struct {
		resActivityLog interface {
			ResourceActivities(ctx context.Context, limit uint, cur string, from *time.Time, to *time.Time) (rsp *feed.Response, err error)
		}
	}
)

func Feed() *feeder {
	return &feeder{
		resActivityLog: feed.ResourceActivity(),
	}
}

func (ctrl feeder) Changes(ctx context.Context, r *request.FeedChanges) (interface{}, error) {
	return ctrl.resActivityLog.ResourceActivities(ctx, r.Limit, r.PageCursor, r.From, r.To)
}
