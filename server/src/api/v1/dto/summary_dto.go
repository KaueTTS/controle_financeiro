package dto

type SummaryResponseDto struct {
	Income  int64 `json:"income"`
	Expense int64 `json:"expense"`
	Balance int64 `json:"balance"`
}
