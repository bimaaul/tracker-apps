package httpservice

import (
	"net/http"

	"github.com/bimaaul/tracker-apps/internals/constant"
	"github.com/bimaaul/tracker-apps/internals/tracker/model/entity"
	"github.com/bimaaul/tracker-apps/internals/tracker/model/request"
	"github.com/bimaaul/tracker-apps/internals/tracker/model/response"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateTransaction(context *fiber.Ctx) error {
	transaction := request.CreateTransactionRequest{}

	err := context.BodyParser(&transaction)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Request failed",
		})
	}

	err = h.DB.Create(&entity.Transaction{
		Amount:          transaction.Amount,
		TransactionType: transaction.TransactionType,
		Notes:           transaction.Notes,
		CreatedAt:       transaction.CreatedAt,
	}).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not create new transaction"},
		)
		return nil
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "New transaction created!"},
	)

	return nil
}

func (h *Handler) DeleteTransaction(context *fiber.Ctx) error {
	transactionModel := entity.Transaction{}

	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "Id cannot be empty!"},
		)
	}

	err := h.DB.Delete(transactionModel, id)
	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not delete transaction"},
		)
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Transaction deleted successfuly!"},
	)

	return nil
}

func (h *Handler) GetTransactionsSummary(context *fiber.Ctx) error {
	type TransactionSummary struct {
		TransactionType constant.TransactionType `json:"transaction_type"`
		TotalAmount     int                      `json:"total_amount"`
	}

	transactionSummary := &[]TransactionSummary{}

	h.DB.Raw(`
	SELECT 
		transaction_type, SUM(t.amount) as total_amount 
	FROM 
		transactions as t
	GROUP BY 
		t.transaction_type;
	`).Scan(&transactionSummary)

	balance := 0
	totalExpense := 0
	totalIncome := 0
	for _, data := range *transactionSummary {
		switch data.TransactionType {
		case constant.INCOME:
			balance += data.TotalAmount
			totalIncome += data.TotalAmount
		case constant.EXPENSE:
			balance -= data.TotalAmount
			totalExpense += data.TotalAmount
		}
	}

	response := response.TransactionSummaryResponse{
		Balance:      balance,
		TotalExpense: totalExpense,
		TotalIncome:  totalIncome,
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"data":    response,
			"message": "Success get transaction summary!",
		},
	)

	return nil
}

func (h *Handler) GetTransactions(context *fiber.Ctx) error {
	transactionModels := &[]entity.Transaction{}

	err := h.DB.Find(transactionModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get transactions data!"},
		)
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "Success get Transactions!",
			"data":    transactionModels,
		},
	)

	return nil
}
