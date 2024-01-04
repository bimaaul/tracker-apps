package repository

import (
	"context"

	"github.com/bimaaul/tracker-apps/internals/cashflow/model/entity"
)

type CashflowRepositoryProvider interface {
	InsertTransaction(ctx context.Context, payload *entity.Transaction) error
	GetTransactions(ctx context.Context) ([]entity.Transaction, error)
	GetTransactionsSummary(ctx context.Context) ([]entity.TransactionSummary, error)
	DeleteTransactionById(ctx context.Context, transactionId int) error
}
