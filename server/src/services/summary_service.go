package services

import (
	"context"
	"controle_financeiro/src/api/v1/dto"
	repository_interfaces "controle_financeiro/src/repositories/interfaces"
)

type SummaryService struct {
	SqliteTransactionRepositoryInterface repository_interfaces.SqliteTransactionRepositoryInterface
}

func NewSummaryService(sqliteTransactionRepositoryInterface repository_interfaces.SqliteTransactionRepositoryInterface) *SummaryService {
	return &SummaryService{
		SqliteTransactionRepositoryInterface: sqliteTransactionRepositoryInterface,
	}
}

func (s *SummaryService) GetSummary(ctx context.Context) (dto.SummaryResponseDto, error) {
	transactions, err := s.SqliteTransactionRepositoryInterface.ListTransactions(ctx, dto.TransactionFilterDto{})
	if err != nil {
		return dto.SummaryResponseDto{}, err
	}

	var income int64
	var expense int64
	var balance int64

	for _, transaction := range transactions {
		if transaction.Amount > 0 {
			income += transaction.Amount
		} else {
			expense += -transaction.Amount
		}

		balance += transaction.Amount
	}

	return dto.SummaryResponseDto{
		Income:  income,
		Expense: expense,
		Balance: balance,
	}, nil
}
