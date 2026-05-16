package services_test

import (
	"context"
	"controle_financeiro/src/api/v1/dto"
	sqlite_mocks "controle_financeiro/src/repositories/sqlite/mocks"
	"controle_financeiro/src/services"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSummary(t *testing.T) {
	t.Run("should return summary", func(t *testing.T) {
		ctx := context.Background()
		expectedSummary := dto.SummaryResponseDto{
			Income:  1000,
			Expense: 500,
			Balance: 500,
		}

		mockSummaryRepository := new(sqlite_mocks.SummaryRepositoryMock)
		mockSummaryRepository.On("GetSummary", ctx).Return(expectedSummary, nil)

		service := services.NewSummaryService(mockSummaryRepository)

		response, err := service.GetSummary(ctx)

		assert.NoError(t, err)
		assert.Equal(t, expectedSummary, response)
		mockSummaryRepository.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		ctx := context.Background()
		expectedError := errors.New("repository error")

		mockSummaryRepository := new(sqlite_mocks.SummaryRepositoryMock)
		mockSummaryRepository.On("GetSummary", ctx).Return(dto.SummaryResponseDto{}, expectedError)

		service := services.NewSummaryService(mockSummaryRepository)

		response, err := service.GetSummary(ctx)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, dto.SummaryResponseDto{}, response)
		mockSummaryRepository.AssertExpectations(t)
	})
}
