package request

import (
	"time"

	"github.com/bimaaul/tracker-apps/internals/constant"
)

type CreateTransactionRequest struct {
	Amount          uint                     `json:"amount"`
	Notes           string                   `json:"notes"`
	TransactionType constant.TransactionType `json:"transaction_type"`
	CreatedAt       time.Time                `json:"created_at"`
}
