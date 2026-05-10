package services_interfaces

import (
	"context"
	"controle_financeiro/src/api/v1/dto"
)

type TransactionServiceInterface interface {
	ListTransactions(ctx context.Context, filters dto.FilterDto) ([]dto.TransactionResponseDto, error)
	CreateTransaction(ctx context.Context, request dto.TransactionRequestDto) error
	DeleteTransaction(ctx context.Context, id uint) error
	UpdateTransaction(ctx context.Context, id uint, request dto.TransactionRequestDto) error
}
