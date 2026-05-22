package services_mocks

import (
	"context"
	dto_summary "controle_financeiro/src/api/v1/dto/summary"

	"github.com/stretchr/testify/mock"
)

type SummaryServiceMock struct {
	mock.Mock
}

func (m *SummaryServiceMock) GetSummary(ctx context.Context, filters dto_summary.SummaryFilterDto) (dto_summary.SummaryResponseDto, error) {
	args := m.Called(ctx, filters)

	return args.Get(0).(dto_summary.SummaryResponseDto), args.Error(1)
}
