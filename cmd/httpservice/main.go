package main

import (
	"log"
	"net/http"

	"github.com/bimaaul/tracker-apps/cmd/internal"
	"github.com/bimaaul/tracker-apps/internals/constant"
	"github.com/bimaaul/tracker-apps/internals/tracker/model/entity"
	"github.com/bimaaul/tracker-apps/internals/tracker/model/request"
	"github.com/bimaaul/tracker-apps/internals/tracker/model/response"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) GetTransactions(context *fiber.Ctx) error {
	transactionModels := &[]entity.Transaction{}

	err := r.DB.Find(transactionModels).Error
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

func (r *Repository) CreateTransaction(context *fiber.Ctx) error {
	transaction := request.CreateTransactionRequest{}

	err := context.BodyParser(&transaction)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Request failed",
		})
	}

	err = r.DB.Create(&entity.Transaction{
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

func (r *Repository) DeleteTransaction(context *fiber.Ctx) error {
	transactionModel := entity.Transaction{}

	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "Id cannot be empty!"},
		)
	}

	err := r.DB.Delete(transactionModel, id)
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

func (r *Repository) GetTransactionsSummary(context *fiber.Ctx) error {
	type TransactionSummary struct {
		TransactionType constant.TransactionType `json:"transaction_type"`
		TotalAmount     int                      `json:"total_amount"`
	}

	transactionSummary := &[]TransactionSummary{}

	r.DB.Raw(`
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

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/transactions", r.GetTransactions)
	api.Get("/transactions/summary", r.GetTransactionsSummary)
	api.Post("/transaction", r.CreateTransaction)
	api.Delete("/transaction/:id", r.DeleteTransaction)
}

func main() {
	// Read Config
	config := internal.InitConfig()

	db, err := internal.NewConnection(&config)
	if err != nil {
		log.Fatal("could not load the database")
	}

	err = entity.MigrateTransaction(db)
	if err != nil {
		log.Fatal("could not migrate database")
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
