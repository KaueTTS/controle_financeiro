package services

import (
	"context"
	"controle_financeiro/src/api/v1/dto"
	repository_interfaces "controle_financeiro/src/repositories/sqlite/interfaces"
)

type SummaryService struct {
	SqliteSummaryRepositoryInterface repository_interfaces.SqliteSummaryRepositoryInterface
}

func NewSummaryService(sqliteSummaryRepositoryInterface repository_interfaces.SqliteSummaryRepositoryInterface) *SummaryService {
	return &SummaryService{
		SqliteSummaryRepositoryInterface: sqliteSummaryRepositoryInterface,
	}
}

func (s *SummaryService) GetSummary(ctx context.Context) (dto.SummaryResponseDto, error) {
	return s.SqliteSummaryRepositoryInterface.GetSummary(ctx)
}
