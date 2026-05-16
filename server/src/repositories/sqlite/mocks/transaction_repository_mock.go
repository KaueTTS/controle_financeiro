package sqlite_mocks

import (
	"context"
	"controle_financeiro/src/api/v1/dto"
	"controle_financeiro/src/models"

	"github.com/stretchr/testify/mock"
)

type TransactionRepositoryMock struct {
	mock.Mock
}

func (m *TransactionRepositoryMock) ListTransactions(ctx context.Context, filters dto.FilterDto) ([]models.Transaction, int64, error) {
	args := m.Called(ctx, filters)

	return args.Get(0).([]models.Transaction), args.Get(1).(int64), args.Error(2)
}

func (m *TransactionRepositoryMock) CreateTransaction(ctx context.Context, transaction models.Transaction) error {
	args := m.Called(ctx, transaction)
	return args.Error(0)
}

func (m *TransactionRepositoryMock) DeleteTransaction(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *TransactionRepositoryMock) UpdateTransaction(ctx context.Context, id uint, transaction models.Transaction) error {
	args := m.Called(ctx, id, transaction)
	return args.Error(0)
}
