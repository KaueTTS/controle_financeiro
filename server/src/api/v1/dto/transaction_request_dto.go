package dto

type TransactionRequestDto struct {
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	Amount      float64 `json:"amount"`
	Category    string  `json:"category"`
	Type        string  `json:"type"`
}
