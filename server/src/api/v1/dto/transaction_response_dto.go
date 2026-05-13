package dto

import "time"

type TransactionDto struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description,omitempty"`
	Amount      float64    `json:"amount"`
	Category    string     `json:"category"`
	Type        string     `json:"type"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt"`
}

type TransactionResponseDto struct {
	Pagination PaginationDto    `json:"pagination"`
	Data       []TransactionDto `json:"data"`
}
