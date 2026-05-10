package services

import (
	"context"
	"controle_financeiro/src/api/v1/dto"
	repository_interfaces "controle_financeiro/src/repositories/sqlite/interfaces"
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
	return s.SqliteTransactionRepositoryInterface.GetSummary(ctx)
}
