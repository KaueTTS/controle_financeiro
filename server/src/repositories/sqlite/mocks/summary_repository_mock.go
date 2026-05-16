package sqlite_mocks

import (
	"context"
	"controle_financeiro/src/api/v1/dto"

	"github.com/stretchr/testify/mock"
)

type SummaryRepositoryMock struct {
	mock.Mock
}

func (m *SummaryRepositoryMock) GetSummary(ctx context.Context) (dto.SummaryResponseDto, error) {
	args := m.Called(ctx)

	return args.Get(0).(dto.SummaryResponseDto), args.Error(1)
}
