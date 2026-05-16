package services_test

import (
	"context"
	dto_summary "controle_financeiro/src/api/v1/dto/summary"
	repositories_mocks "controle_financeiro/src/mocks/repositories"
	services "controle_financeiro/src/services"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSummary(t *testing.T) {
	t.Run("should return summary", func(t *testing.T) {
		ctx := context.Background()
		expectedSummary := dto_summary.SummaryResponseDto{
			Income:  1000,
			Expense: 500,
			Balance: 500,
		}

		mockSummaryRepository := new(repositories_mocks.SummaryRepositoryMock)
		mockSummaryRepository.On("GetSummary", ctx, dto_summary.SummaryFilterDto{}).Return(expectedSummary, nil)

		service := services.NewSummaryService(mockSummaryRepository)

		response, err := service.GetSummary(ctx, dto_summary.SummaryFilterDto{})

		assert.NoError(t, err)
		assert.Equal(t, expectedSummary, response)
		mockSummaryRepository.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		ctx := context.Background()
		expectedError := errors.New("repository error")

		mockSummaryRepository := new(repositories_mocks.SummaryRepositoryMock)
		mockSummaryRepository.On("GetSummary", ctx, dto_summary.SummaryFilterDto{}).Return(dto_summary.SummaryResponseDto{}, expectedError)

		service := services.NewSummaryService(mockSummaryRepository)

		response, err := service.GetSummary(ctx, dto_summary.SummaryFilterDto{})

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, dto_summary.SummaryResponseDto{}, response)
		mockSummaryRepository.AssertExpectations(t)
	})
}
