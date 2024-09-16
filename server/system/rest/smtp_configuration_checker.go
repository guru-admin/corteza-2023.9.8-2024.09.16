package rest

import (
	"context"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/rest/request"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/service"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
)

type (
	SmtpConfigurationChecker struct {
		svc smtpConfigurationCheckerService
	}

	smtpConfigurationCheckerService interface {
		Check(context.Context, *types.SmtpConfiguration) (*types.SmtpCheckResult, error)
	}
)

func (SmtpConfigurationChecker) New() *SmtpConfigurationChecker {
	return &SmtpConfigurationChecker{
		svc: service.DefaultSMTPChecker,
	}
}

func (ctrl *SmtpConfigurationChecker) Check(ctx context.Context, r *request.SmtpConfigurationCheckerCheck) (interface{}, error) {
	var (
		err          error
		checkResults = &types.SmtpCheckResult{}
		smtp         = &types.SmtpConfiguration{
			Host:          r.Host,
			Port:          r.Port,
			Recipients:    r.Recipients,
			Username:      r.Username,
			Password:      r.Password,
			TLSInsecure:   r.TlsInsecure,
			TLSServerName: r.TlsServerName,
		}
	)

	checkResults, err = ctrl.svc.Check(ctx, smtp)
	if err != nil {
		return nil, err
	}

	return checkResults, nil
}
