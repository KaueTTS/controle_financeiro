package services_mocks

import (
	"context"
	dto_transaction "controle_financeiro/src/api/v1/dto/transaction"

	"github.com/stretchr/testify/mock"
)

type TransactionServiceMock struct {
	mock.Mock
}

func (m *TransactionServiceMock) ListTransactions(ctx context.Context, filters dto_transaction.TransactionFilterDto) (dto_transaction.TransactionResponseDto, error) {
	args := m.Called(ctx, filters)

	return args.Get(0).(dto_transaction.TransactionResponseDto), args.Error(1)
}

func (m *TransactionServiceMock) CreateTransaction(ctx context.Context, request dto_transaction.TransactionRequestDto) error {
	args := m.Called(ctx, request)
	return args.Error(0)
}

func (m *TransactionServiceMock) DeleteTransaction(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *TransactionServiceMock) UpdateTransaction(ctx context.Context, id uint, request dto_transaction.TransactionRequestDto) error {
	args := m.Called(ctx, id, request)
	return args.Error(0)
}
