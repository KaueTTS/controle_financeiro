package services_interfaces

import (
	"context"
	dto_transaction "controle_financeiro/src/api/v1/dto/transaction"
)

type TransactionServiceInterface interface {
	ListTransactions(ctx context.Context, filters dto_transaction.TransactionFilterDto) (dto_transaction.TransactionResponseDto, error)
	CreateTransaction(ctx context.Context, request dto_transaction.TransactionRequestDto) error
	DeleteTransaction(ctx context.Context, id uint) error
	UpdateTransaction(ctx context.Context, id uint, request dto_transaction.TransactionRequestDto) error
}
