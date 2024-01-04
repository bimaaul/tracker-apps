package main

import (
	cashflow "github.com/bimaaul/tracker-apps/internals/cashflow/httpservice"
)

type HTTPService struct {
	Cashflow *cashflow.Handler
}
