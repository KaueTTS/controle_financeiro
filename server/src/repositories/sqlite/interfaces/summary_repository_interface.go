package repository_interfaces

import (
	"context"
	"controle_financeiro/src/api/v1/dto"
)

type SqliteSummaryRepositoryInterface interface {
	GetSummary(ctx context.Context) (dto.SummaryResponseDto, error)
}
