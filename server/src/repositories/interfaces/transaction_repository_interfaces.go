package repository_interfaces

import (
	"context"
	"controle_financeiro/src/api/v1/dto"
	"controle_financeiro/src/models"
)

type SqliteTransactionRepositoryInterface interface {
	ListTransactions(ctx context.Context, filters dto.TransactionFilterDto) ([]models.Transaction, error)
	CreateTransaction(ctx context.Context, transaction models.Transaction) error
	DeleteTransaction(ctx context.Context, id uint) error
}
