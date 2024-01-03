package service

import (
	"context"
)

type TrackerServiceProvider interface {
	GetTransactions(ctx context.Context) error
	CreateTransaction(ctx context.Context) error
	DeleteTransaction(ctx context.Context) error
	GetTransactionsSummary(ctx context.Context) error
}
