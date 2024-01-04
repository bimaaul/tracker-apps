package internal

import cashflowService "github.com/bimaaul/tracker-apps/internals/cashflow/service"

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

type Service struct {
	Cashflow cashflowService.CashflowServiceProvider
}
