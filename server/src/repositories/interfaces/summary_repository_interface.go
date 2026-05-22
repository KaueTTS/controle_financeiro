package repository_interfaces

import (
	"context"
	dto_summary "controle_financeiro/src/api/v1/dto/summary"
)

type SummaryRepositoryInterface interface {
	GetSummary(ctx context.Context, filters dto_summary.SummaryFilterDto) (dto_summary.SummaryResponseDto, error)
}
