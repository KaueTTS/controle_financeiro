package dto

import "time"

type TransactionResponseDto struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description,omitempty"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
}
