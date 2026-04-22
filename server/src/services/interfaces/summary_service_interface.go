package services_interfaces

import (
	"context"
	"controle_financeiro/src/api/v1/dto"
)

type SummaryServiceInterface interface {
	GetSummary(ctx context.Context) (dto.SummaryResponseDto, error)
}
