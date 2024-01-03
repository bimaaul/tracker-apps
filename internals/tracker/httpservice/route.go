package httpservice

import "github.com/gofiber/fiber/v2"

func (h *Handler) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/transactions", h.CreateTransaction)
	api.Get("/transactions/summary", h.GetTransactionsSummary)
	api.Post("/transaction", h.CreateTransaction)
	api.Delete("/transaction/:id", h.DeleteTransaction)
}
