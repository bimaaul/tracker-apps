package entity

import (
	"time"

	"github.com/bimaaul/tracker-apps/internals/constant"
	"gorm.io/gorm"
)

type Transaction struct {
	ID              uint                     `gorm:"primary key;autoIncrement" json:"id"`
	Amount          uint                     `json:"amount"`
	Notes           string                   `json:"notes"`
	TransactionType constant.TransactionType `gorm:"type:transaction_type" json:"transaction_type"`
	CreatedAt       time.Time                `json:"created_at"`
}

func MigrateTransaction(db *gorm.DB) error {
	db.Exec(`
	do $$ BEGIN
		CREATE TYPE transaction_type AS ENUM ('income', 'expense');
	EXCEPTION
		WHEN duplicate_object THEN null;
	end $$;
	`)
	err := db.AutoMigrate(&Transaction{})
	return err
}
