package services

import (
	"context"
	dto_summary "controle_financeiro/src/api/v1/dto/summary"
	repository_interfaces "controle_financeiro/src/repositories/interfaces"
)

type SummaryService struct {
	SummaryRepositoryInterface repository_interfaces.SummaryRepositoryInterface
}

func NewSummaryService(summaryRepositoryInterface repository_interfaces.SummaryRepositoryInterface) *SummaryService {
	return &SummaryService{
		SummaryRepositoryInterface: summaryRepositoryInterface,
	}
}

func (s *SummaryService) GetSummary(ctx context.Context, filters dto_summary.SummaryFilterDto) (dto_summary.SummaryResponseDto, error) {
	return s.SummaryRepositoryInterface.GetSummary(ctx, filters)
}
