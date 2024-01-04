package httpservice

import (
	"net/http"
	"strconv"

	"github.com/bimaaul/tracker-apps/internals/cashflow/model/request"
	"github.com/gofiber/fiber/v2"
)

func NewHandler(cfg HandlerConfig) *Handler {
	return &Handler{
		cashflowSrv: cfg.CashflowSrv,
	}
}

func (h *Handler) CreateTransaction(context *fiber.Ctx) error {
	request := new(request.CreateTransactionRequest)

	err := context.BodyParser(request)
	if err != nil {
		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Request failed",
		})
	}

	err = h.cashflowSrv.CreateTransaction(context.Context(), *request)
	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not create new transaction"},
		)

	}

	return context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "New transaction created!"},
	)

}

func (h *Handler) DeleteTransaction(context *fiber.Ctx) error {

	transactionId, errParse := strconv.Atoi(context.Params("id"))
	if errParse != nil {
		return context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "Id is not valid!"},
		)
	}

	err := h.cashflowSrv.DeleteTransaction(context.Context(), transactionId)

	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not delete transaction"},
		)
	}

	return context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Transaction deleted successfuly!"},
	)

}

func (h *Handler) GetTransactionsSummary(context *fiber.Ctx) error {
	txSummary, err := h.cashflowSrv.GetTransactionsSummary(context.Context())

	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not get transaction summary"},
		)

	}

	return context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"data":    txSummary,
			"message": "Success get transaction summary!",
		},
	)
}

func (h *Handler) GetTransactions(context *fiber.Ctx) error {
	data, err := h.cashflowSrv.GetTransactions(context.Context())

	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get transactions data!"},
		)
	}

	return context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "Success get Transactions!",
			"data":    data,
		},
	)
}
