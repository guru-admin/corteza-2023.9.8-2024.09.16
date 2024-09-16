package rest

import (
	"context"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/rest/request"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/service"
)

type (
	Stats struct {
		svc statsService
	}

	statsService interface {
		Metrics(context.Context) (*service.StatisticsMetricsPayload, error)
	}
)

func (Stats) New() *Stats {
	return &Stats{
		svc: service.DefaultStatistics,
	}
}

func (ctrl *Stats) List(ctx context.Context, r *request.StatsList) (interface{}, error) {
	return ctrl.svc.Metrics(ctx)
}
