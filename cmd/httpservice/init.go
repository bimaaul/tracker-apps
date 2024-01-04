package main

import (
	"github.com/bimaaul/tracker-apps/cmd/internal"
	cashflow "github.com/bimaaul/tracker-apps/internals/cashflow/httpservice"
)

func InitializeService(serv *internal.Service) HTTPService {
	return HTTPService{
		Cashflow: cashflow.NewHandler(cashflow.HandlerConfig{
			CashflowSrv: serv.Cashflow,
		}),
	}
}
