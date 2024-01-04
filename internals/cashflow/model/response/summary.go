package response

type TransactionSummaryResponse struct {
	Balance      int `json:"amount"`
	TotalExpense int `json:"total_expense"`
	TotalIncome  int `json:"total_income"`
}
