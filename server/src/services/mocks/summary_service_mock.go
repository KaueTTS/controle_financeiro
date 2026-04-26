package services_mocks

import (
	"context"
	"controle_financeiro/src/api/v1/dto"

	"github.com/stretchr/testify/mock"
)

type SummaryServiceMock struct {
	mock.Mock
}

func (m *SummaryServiceMock) GetSummary(ctx context.Context) (dto.SummaryResponseDto, error) {
	args := m.Called(ctx)

	return args.Get(0).(dto.SummaryResponseDto), args.Error(1)
}
