package repository_interfaces

import (
	"context"
	dto_transaction "controle_financeiro/src/api/v1/dto/transaction"
	"controle_financeiro/src/models"
)

type TransactionRepositoryInterface interface {
	ListTransactions(ctx context.Context, filters dto_transaction.TransactionFilterDto) ([]models.Transaction, int64, error)
	CreateTransaction(ctx context.Context, transaction models.Transaction) error
	DeleteTransaction(ctx context.Context, id uint) error
	UpdateTransaction(ctx context.Context, id uint, transaction models.Transaction) error
}
