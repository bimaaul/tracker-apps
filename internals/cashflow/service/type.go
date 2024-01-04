package service

import (
	"context"

	"github.com/bimaaul/tracker-apps/internals/cashflow/model/entity"
	"github.com/bimaaul/tracker-apps/internals/cashflow/model/request"
	"github.com/bimaaul/tracker-apps/internals/cashflow/model/response"
)

type CashflowServiceProvider interface {
	CreateTransaction(ctx context.Context, payload request.CreateTransactionRequest) error
	GetTransactions(ctx context.Context) ([]entity.Transaction, error)
	DeleteTransaction(ctx context.Context, transactionId int) error
	GetTransactionsSummary(ctx context.Context) (response.TransactionSummaryResponse, error)
}
