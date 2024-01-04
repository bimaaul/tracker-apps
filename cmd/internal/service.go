package internal

import (
	cashflowRepo "github.com/bimaaul/tracker-apps/internals/cashflow/repository/postgre"
	cashflowService "github.com/bimaaul/tracker-apps/internals/cashflow/service"
)

func GetService(config Config) *Service {
	// Connect to DB
	db := NewConnection(&config)

	cashflowRepo := cashflowRepo.NewCashflow(db)
	cashflowServ := cashflowService.NewCashflowService(cashflowService.CashflowServiceConfig{
		CashflowServiceRepository: cashflowRepo,
	})

	return &Service{
		Cashflow: cashflowServ,
	}
}
