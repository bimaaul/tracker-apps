package service

import (
	"context"

	"github.com/bimaaul/tracker-apps/internals/cashflow/model/entity"
	"github.com/bimaaul/tracker-apps/internals/cashflow/model/request"
	"github.com/bimaaul/tracker-apps/internals/cashflow/model/response"
	"github.com/bimaaul/tracker-apps/internals/cashflow/repository"
	"github.com/bimaaul/tracker-apps/internals/constant"
)

type cashflowService struct {
	cashflowRepository repository.CashflowRepositoryProvider
}

type CashflowServiceConfig struct {
	CashflowServiceRepository repository.CashflowRepositoryProvider
}

func NewCashflowService(config CashflowServiceConfig) CashflowServiceProvider {
	return &cashflowService{
		cashflowRepository: config.CashflowServiceRepository,
	}

}

func (c *cashflowService) CreateTransaction(ctx context.Context, payload request.CreateTransactionRequest) error {
	err := c.cashflowRepository.InsertTransaction(ctx, &entity.Transaction{
		Amount:          payload.Amount,
		TransactionType: payload.TransactionType,
		Notes:           payload.Notes,
		CreatedAt:       payload.CreatedAt,
	})

	if err != nil {
		return err
	}

	return nil
}

func (c *cashflowService) GetTransactions(ctx context.Context) ([]entity.Transaction, error) {
	data, err := c.cashflowRepository.GetTransactions(ctx)

	if err != nil {
		return data, err
	}

	return data, nil
}

func (c *cashflowService) GetTransactionsSummary(ctx context.Context) (response.TransactionSummaryResponse, error) {
	response := response.TransactionSummaryResponse{}

	txSummary, err := c.cashflowRepository.GetTransactionsSummary(ctx)
	if err != nil {
		return response, err
	}

	for _, data := range txSummary {
		switch data.TransactionType {
		case constant.INCOME:
			response.Balance += data.TotalAmount
			response.TotalIncome += data.TotalAmount
		case constant.EXPENSE:
			response.Balance -= data.TotalAmount
			response.TotalExpense += data.TotalAmount
		}
	}

	return response, nil
}

func (c *cashflowService) DeleteTransaction(ctx context.Context, transactionId int) error {
	err := c.cashflowRepository.DeleteTransactionById(ctx, transactionId)

	if err != nil {
		return err
	}

	return nil
}
