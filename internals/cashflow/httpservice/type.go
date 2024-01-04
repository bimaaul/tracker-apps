package httpservice

import (
	service "github.com/bimaaul/tracker-apps/internals/cashflow/service"
	"gorm.io/gorm"
)

type Handler struct {
	DB          *gorm.DB
	cashflowSrv service.CashflowServiceProvider
}

type HandlerConfig struct {
	CashflowSrv service.CashflowServiceProvider
}
