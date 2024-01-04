package postgre

import (
	"context"

	"github.com/bimaaul/tracker-apps/internals/cashflow/model/entity"
	"github.com/bimaaul/tracker-apps/internals/cashflow/repository"

	"gorm.io/gorm"
)

type cashflowRepo struct {
	db *gorm.DB
}

func NewCashflow(db *gorm.DB) repository.CashflowRepositoryProvider {
	return &cashflowRepo{
		db: db,
	}
}

func (c *cashflowRepo) InsertTransaction(ctx context.Context, payload *entity.Transaction) error {
	err := c.db.Create(payload).Error

	if err != nil {
		return err
	}

	return nil
}

func (c *cashflowRepo) GetTransactions(ctx context.Context) ([]entity.Transaction, error) {
	transactionModels := &[]entity.Transaction{}

	err := c.db.Find(&transactionModels).Error
	if err != nil {
		return *transactionModels, err
	}

	return *transactionModels, nil
}

func (c *cashflowRepo) GetTransactionsSummary(ctx context.Context) ([]entity.TransactionSummary, error) {
	transactionByTotalAmount := &[]entity.TransactionSummary{}

	err := c.db.Raw(`
	SELECT 
		transaction_type, SUM(t.amount) as total_amount 
	FROM 
		transactions as t
	GROUP BY 
		t.transaction_type;
	`).Scan(&transactionByTotalAmount).Error

	if err != nil {
		return *transactionByTotalAmount, err
	}

	return *transactionByTotalAmount, nil
}

func (c *cashflowRepo) DeleteTransactionById(ctx context.Context, transactionId int) error {
	transactionModel := entity.Transaction{}

	err := c.db.Delete(&transactionModel, transactionId).Error
	if err != nil {
		return err
	}

	return nil
}
