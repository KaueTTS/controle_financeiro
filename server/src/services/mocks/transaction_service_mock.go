package services_mocks

import (
	"context"
	"controle_financeiro/src/api/v1/dto"

	"github.com/stretchr/testify/mock"
)

type TransactionServiceMock struct {
	mock.Mock
}

func (m *TransactionServiceMock) ListTransactions(ctx context.Context, filters dto.TransactionFilterDto) ([]dto.TransactionResponseDto, error) {
	args := m.Called(ctx, filters)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]dto.TransactionResponseDto), args.Error(1)
}

func (m *TransactionServiceMock) CreateTransaction(ctx context.Context, request dto.TransactionRequestDto) error {
	args := m.Called(ctx, request)
	return args.Error(0)
}

func (m *TransactionServiceMock) DeleteTransaction(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *TransactionServiceMock) UpdateTransaction(ctx context.Context, id uint, request dto.TransactionRequestDto) error {
	args := m.Called(ctx, id, request)
	return args.Error(0)
}
