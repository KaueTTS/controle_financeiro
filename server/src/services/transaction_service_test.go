package services_test

import (
	"context"
	"controle_financeiro/src/api/v1/dto"
	"controle_financeiro/src/models"
	sqlite_mocks "controle_financeiro/src/repositories/sqlite/mocks"
	"controle_financeiro/src/services"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListTransactions(t *testing.T) {
	t.Run("should return transactions with pagination", func(t *testing.T) {
		ctx := context.Background()
		createdAt := time.Date(2026, 5, 13, 10, 0, 0, 0, time.UTC)
		description := "Pagamento mensal"

		filters := dto.FilterDto{
			Page:    1,
			PerPage: 10,
		}

		transactions := []models.Transaction{
			{
				ID:          1,
				Title:       "Salário",
				Description: &description,
				Amount:      5000,
				Type:        "income",
				Category:    "Emprego",
				CreatedAt:   createdAt,
			},
			{
				ID:        2,
				Title:     "Mercado",
				Amount:    200,
				Type:      "expense",
				Category:  "Alimentação",
				CreatedAt: createdAt,
			},
		}

		mockRepository := new(sqlite_mocks.TransactionRepositoryMock)
		mockRepository.On("ListTransactions", ctx, filters).Return(transactions, int64(2), nil)

		service := services.NewTransactionService(mockRepository)

		response, err := service.ListTransactions(ctx, filters)

		assert.NoError(t, err)
		assert.Equal(t, int64(2), response.Pagination.Total)
		assert.Equal(t, 1, response.Pagination.Page)
		assert.Equal(t, 10, response.Pagination.PerPage)
		assert.Equal(t, 1, response.Pagination.PageCount)
		assert.Len(t, response.Data, 2)
		assert.Equal(t, uint(1), response.Data[0].ID)
		assert.Equal(t, "Salário", response.Data[0].Title)
		assert.Equal(t, &description, response.Data[0].Description)
		assert.Equal(t, float64(5000), response.Data[0].Amount)
		assert.Equal(t, "income", response.Data[0].Type)
		assert.Equal(t, "Emprego", response.Data[0].Category)

		mockRepository.AssertExpectations(t)
	})

	t.Run("should use default pagination when page and perPage are invalid", func(t *testing.T) {
		ctx := context.Background()

		inputFilters := dto.FilterDto{
			Page:    0,
			PerPage: 0,
		}

		expectedFilters := dto.FilterDto{
			Page:    1,
			PerPage: 10,
		}

		mockRepository := new(sqlite_mocks.TransactionRepositoryMock)
		mockRepository.On("ListTransactions", ctx, expectedFilters).Return([]models.Transaction{}, int64(0), nil)

		service := services.NewTransactionService(mockRepository)

		response, err := service.ListTransactions(ctx, inputFilters)

		assert.NoError(t, err)
		assert.Equal(t, 1, response.Pagination.Page)
		assert.Equal(t, 10, response.Pagination.PerPage)
		assert.Equal(t, 0, response.Pagination.PageCount)
		assert.Equal(t, int64(0), response.Pagination.Total)
		assert.Empty(t, response.Data)

		mockRepository.AssertExpectations(t)
	})

	t.Run("should limit perPage to 100 when perPage is greater than 100", func(t *testing.T) {
		ctx := context.Background()

		inputFilters := dto.FilterDto{
			Page:    1,
			PerPage: 200,
		}

		expectedFilters := dto.FilterDto{
			Page:    1,
			PerPage: 100,
		}

		mockRepository := new(sqlite_mocks.TransactionRepositoryMock)
		mockRepository.On("ListTransactions", ctx, expectedFilters).Return([]models.Transaction{}, int64(0), nil)

		service := services.NewTransactionService(mockRepository)

		response, err := service.ListTransactions(ctx, inputFilters)

		assert.NoError(t, err)
		assert.Equal(t, 1, response.Pagination.Page)
		assert.Equal(t, 100, response.Pagination.PerPage)

		mockRepository.AssertExpectations(t)
	})

	t.Run("should calculate pageCount correctly when total is not divisible by perPage", func(t *testing.T) {
		ctx := context.Background()

		filters := dto.FilterDto{
			Page:    1,
			PerPage: 10,
		}

		mockRepository := new(sqlite_mocks.TransactionRepositoryMock)
		mockRepository.On("ListTransactions", ctx, filters).Return([]models.Transaction{}, int64(25), nil)

		service := services.NewTransactionService(mockRepository)

		response, err := service.ListTransactions(ctx, filters)

		assert.NoError(t, err)
		assert.Equal(t, 3, response.Pagination.PageCount)
		assert.Equal(t, int64(25), response.Pagination.Total)

		mockRepository.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		ctx := context.Background()
		expectedError := errors.New("repository error")

		filters := dto.FilterDto{
			Page:    1,
			PerPage: 10,
		}

		mockRepository := new(sqlite_mocks.TransactionRepositoryMock)
		mockRepository.On("ListTransactions", ctx, filters).Return([]models.Transaction{}, int64(0), expectedError)

		service := services.NewTransactionService(mockRepository)

		response, err := service.ListTransactions(ctx, filters)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, dto.TransactionResponseDto{}, response)

		mockRepository.AssertExpectations(t)
	})
}

func TestCreateTransaction(t *testing.T) {
	t.Run("should create transaction", func(t *testing.T) {
		ctx := context.Background()
		description := "Pagamento mensal"

		request := dto.TransactionRequestDto{
			Title:       "Salário",
			Description: &description,
			Amount:      5000,
			Type:        "income",
			Category:    "Emprego",
		}

		expectedTransaction := models.Transaction{
			Title:       "Salário",
			Description: &description,
			Amount:      5000,
			Type:        "income",
			Category:    "Emprego",
		}

		mockRepository := new(sqlite_mocks.TransactionRepositoryMock)
		mockRepository.On("CreateTransaction", ctx, expectedTransaction).Return(nil)

		service := services.NewTransactionService(mockRepository)

		err := service.CreateTransaction(ctx, request)

		assert.NoError(t, err)
		mockRepository.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		ctx := context.Background()
		expectedError := errors.New("repository error")

		request := dto.TransactionRequestDto{
			Title:    "Mercado",
			Amount:   200,
			Type:     "expense",
			Category: "Alimentação",
		}

		expectedTransaction := models.Transaction{
			Title:    "Mercado",
			Amount:   200,
			Type:     "expense",
			Category: "Alimentação",
		}

		mockRepository := new(sqlite_mocks.TransactionRepositoryMock)
		mockRepository.On("CreateTransaction", ctx, expectedTransaction).Return(expectedError)

		service := services.NewTransactionService(mockRepository)

		err := service.CreateTransaction(ctx, request)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		mockRepository.AssertExpectations(t)
	})
}

func TestDeleteTransaction(t *testing.T) {
	t.Run("should delete transaction", func(t *testing.T) {
		ctx := context.Background()
		id := uint(1)

		mockRepository := new(sqlite_mocks.TransactionRepositoryMock)
		mockRepository.On("DeleteTransaction", ctx, id).Return(nil)

		service := services.NewTransactionService(mockRepository)

		err := service.DeleteTransaction(ctx, id)

		assert.NoError(t, err)
		mockRepository.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		ctx := context.Background()
		id := uint(1)
		expectedError := errors.New("repository error")

		mockRepository := new(sqlite_mocks.TransactionRepositoryMock)
		mockRepository.On("DeleteTransaction", ctx, id).Return(expectedError)

		service := services.NewTransactionService(mockRepository)

		err := service.DeleteTransaction(ctx, id)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		mockRepository.AssertExpectations(t)
	})
}

func TestUpdateTransaction(t *testing.T) {
	t.Run("should update transaction", func(t *testing.T) {
		ctx := context.Background()
		id := uint(1)
		description := "Compra do mês"

		request := dto.TransactionRequestDto{
			Title:       "Mercado",
			Description: &description,
			Amount:      300,
			Type:        "expense",
			Category:    "Alimentação",
		}

		expectedTransaction := models.Transaction{
			Title:       "Mercado",
			Description: &description,
			Amount:      300,
			Type:        "expense",
			Category:    "Alimentação",
		}

		mockRepository := new(sqlite_mocks.TransactionRepositoryMock)
		mockRepository.On("UpdateTransaction", ctx, id, expectedTransaction).Return(nil)

		service := services.NewTransactionService(mockRepository)

		err := service.UpdateTransaction(ctx, id, request)

		assert.NoError(t, err)
		mockRepository.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		ctx := context.Background()
		id := uint(1)
		expectedError := errors.New("repository error")

		request := dto.TransactionRequestDto{
			Title:    "Mercado",
			Amount:   300,
			Type:     "expense",
			Category: "Alimentação",
		}

		expectedTransaction := models.Transaction{
			Title:    "Mercado",
			Amount:   300,
			Type:     "expense",
			Category: "Alimentação",
		}

		mockRepository := new(sqlite_mocks.TransactionRepositoryMock)
		mockRepository.On("UpdateTransaction", ctx, id, expectedTransaction).Return(expectedError)

		service := services.NewTransactionService(mockRepository)

		err := service.UpdateTransaction(ctx, id, request)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		mockRepository.AssertExpectations(t)
	})
}
