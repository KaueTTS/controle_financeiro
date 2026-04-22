package dto

import "time"

type TransactionRequestDto struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Category    string  `json:"category"`
	Type        string  `json:"type"`
}

type TransactionResponseDto struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
}

type TransactionFilterDto struct {
	Search   string `query:"search"`
	Type     string `query:"type"`
	Category string `query:"category"`
}
