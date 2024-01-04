package entity

import "github.com/bimaaul/tracker-apps/internals/constant"

type TransactionSummary struct {
	TransactionType constant.TransactionType `json:"transaction_type"`
	TotalAmount     int                      `json:"total_amount"`
}
