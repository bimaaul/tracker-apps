package main

import (
	"github.com/bimaaul/tracker-apps/cmd/internal"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Read Config
	config := internal.InitConfig()

	// Get all available services
	serv := internal.GetService(config)

	httpService := InitializeService(serv)

	app := fiber.New()
	httpService.Cashflow.SetupRoutes(app)

	app.Listen(":8080")
}
