package main

import (
	"log"

	"github.com/bimaaul/tracker-apps/cmd/internal"
	trackerRepository "github.com/bimaaul/tracker-apps/internals/tracker/httpservice"
	"github.com/bimaaul/tracker-apps/internals/tracker/model/entity"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Read Config
	config := internal.InitConfig()

	// Connect to DB
	db, err := internal.NewConnection(&config)
	if err != nil {
		log.Fatal("could not load the database")
	}

	// Run migration
	err = entity.MigrateTransaction(db)
	if err != nil {
		log.Fatal("could not migrate database")
	}

	trackerRepository := trackerRepository.Handler{
		DB: db,
	}
	app := fiber.New()

	trackerRepository.SetupRoutes(app)

	app.Listen(":8080")
}
