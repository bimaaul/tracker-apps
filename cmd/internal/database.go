package internal

import (
	"fmt"
	"log"

	"github.com/bimaaul/tracker-apps/internals/cashflow/model/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(config *Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("could not load the database")
	}

	err = entity.MigrateTransaction(db)
	if err != nil {
		log.Fatal("could not migrate database")
	}

	return db
}
