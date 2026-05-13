package controllers_test

import (
	"controle_financeiro/src/api/v1/controllers"
	"controle_financeiro/src/api/v1/dto"
	services_mocks "controle_financeiro/src/services/mocks"
	utils_errors "controle_financeiro/src/utils/errors"
	resolvers "controle_financeiro/src/utils/resolvers"
	"errors"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListTransactions(t *testing.T) {
	t.Run("should return list of transactions with status 200", func(t *testing.T) {
		app := fiber.New()

		mockTransactionService := new(services_mocks.TransactionServiceMock)
		mockTransactionService.On("ListTransactions", mock.Anything, mock.Anything).Return(dto.TransactionResponseDto{
			Pagination: dto.PaginationDto{
				Page:      1,
				PerPage:   10,
				PageCount: 1,
				Total:     2,
			},
			Data: []dto.TransactionDto{
				{
					ID:        1,
					Title:     "Salary",
					Amount:    5000,
					Type:      "income",
					Category:  "Emprego",
					CreatedAt: time.Time{},
					UpdatedAt: nil,
					DeletedAt: nil,
				},
				{
					ID:        2,
					Title:     "Groceries",
					Amount:    200,
					Type:      "expense",
					Category:  "Alimentação",
					CreatedAt: time.Time{},
					UpdatedAt: nil,
					DeletedAt: nil,
				},
			},
		}, nil)

		controller := controllers.NewTransactionController(mockTransactionService)
		app.Get("/transactions", controller.ListTransactions)

		req := httptest.NewRequest("GET", "/transactions", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.JSONEq(t, `{
			"pagination": {
				"page": 1,
				"perPage": 10,
				"pageCount": 1,
				"total": 2
			},
			"data": [
				{
					"id": 1,
					"title": "Salary",
					"amount": 5000,
					"type": "income",
					"category": "Emprego",
					"createdAt": "0001-01-01T00:00:00Z",
					"updatedAt": null,
					"deletedAt": null
				},
				{
					"id": 2,
					"title": "Groceries",
					"amount": 200,
					"type": "expense",
					"category": "Alimentação",
					"createdAt": "0001-01-01T00:00:00Z",
					"updatedAt": null,
					"deletedAt": null
				}
			]
		}`, string(body))

		mockTransactionService.AssertExpectations(t)
	})

	t.Run("should return error with status 500 when service fails", func(t *testing.T) {
		app := fiber.New()

		mockTransactionService := new(services_mocks.TransactionServiceMock)
		mockTransactionService.On("ListTransactions", mock.Anything, mock.Anything).
			Return(dto.TransactionResponseDto{
				Data:       []dto.TransactionDto{},
				Pagination: dto.PaginationDto{},
			}, errors.New("internal error"))

		controller := controllers.NewTransactionController(mockTransactionService)
		app.Get("/transactions", controller.ListTransactions)

		req := httptest.NewRequest("GET", "/transactions", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.JSONEq(t, `{
			"error": "internal error"
		}`, string(body))

		mockTransactionService.AssertExpectations(t)
	})
}

func TestCreateTransaction(t *testing.T) {
	t.Run("should create transaction with status 201", func(t *testing.T) {
		app := fiber.New()

		mockTransactionService := new(services_mocks.TransactionServiceMock)
		mockTransactionService.On("CreateTransaction", mock.Anything, mock.Anything).Return(nil)

		controller := controllers.NewTransactionController(mockTransactionService)
		app.Post("/transactions", controller.CreateTransaction)

		requestBody := `{
			"title": "Salário",
			"amount": 5000,
			"type": "income",
			"category": "Emprego"
		}`

		req := httptest.NewRequest("POST", "/transactions", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.JSONEq(t, `{
			"message": "`+resolvers.TransactionCreated+`"
		}`, string(body))

		mockTransactionService.AssertExpectations(t)
	})

	t.Run("should return error with status 400 when title is empty", func(t *testing.T) {
		app := fiber.New()

		mockTransactionService := new(services_mocks.TransactionServiceMock)

		controller := controllers.NewTransactionController(mockTransactionService)
		app.Post("/transactions", controller.CreateTransaction)

		requestBody := `{
			"title": "",
			"amount": 5000,
			"type": "income",
			"category": "Emprego"
		}`

		req := httptest.NewRequest("POST", "/transactions", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.JSONEq(t, `{
			"error": "`+utils_errors.TitleRequired+`"
		}`, string(body))

		mockTransactionService.AssertNotCalled(t, "CreateTransaction")
	})

	t.Run("should return error with status 400 when amount is invalid", func(t *testing.T) {
		app := fiber.New()

		mockTransactionService := new(services_mocks.TransactionServiceMock)

		controller := controllers.NewTransactionController(mockTransactionService)
		app.Post("/transactions", controller.CreateTransaction)

		requestBody := `{
			"title": "Salary",
			"amount": 0,
			"category": "Job",
			"type": "income"
		}`

		req := httptest.NewRequest("POST", "/transactions", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.JSONEq(t, `{
			"error": "`+utils_errors.AmountRequired+`"
		}`, string(body))

		mockTransactionService.AssertNotCalled(t, "CreateTransaction")
	})

	t.Run("should return error with status 400 when category is empty", func(t *testing.T) {
		app := fiber.New()

		mockTransactionService := new(services_mocks.TransactionServiceMock)

		controller := controllers.NewTransactionController(mockTransactionService)
		app.Post("/transactions", controller.CreateTransaction)

		requestBody := `{
			"title": "Salary",
			"amount": 5000,
			"category": "",
			"type": "income"
		}`

		req := httptest.NewRequest("POST", "/transactions", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.JSONEq(t, `{
			"error": "`+utils_errors.CategoryRequired+`"
		}`, string(body))

		mockTransactionService.AssertNotCalled(t, "CreateTransaction")
	})

	t.Run("should return error with status 400 when type is invalid", func(t *testing.T) {
		app := fiber.New()

		mockTransactionService := new(services_mocks.TransactionServiceMock)

		controller := controllers.NewTransactionController(mockTransactionService)
		app.Post("/transactions", controller.CreateTransaction)

		requestBody := `{
			"title": "Salary",
			"amount": 5000,
			"category": "Job",
			"type": "invalid"
		}`

		req := httptest.NewRequest("POST", "/transactions", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.JSONEq(t, `{
			"error": "`+utils_errors.TypeInvalid+`"
		}`, string(body))

		mockTransactionService.AssertNotCalled(t, "CreateTransaction")
	})

	t.Run("should return error with status 500 when service fails", func(t *testing.T) {
		app := fiber.New()

		mockTransactionService := new(services_mocks.TransactionServiceMock)
		mockTransactionService.On("CreateTransaction", mock.Anything, mock.Anything).
			Return(errors.New("internal error"))

		controller := controllers.NewTransactionController(mockTransactionService)
		app.Post("/transactions", controller.CreateTransaction)

		requestBody := `{
			"title": "Salary",
			"amount": 5000,
			"category": "Job",
			"type": "income"
		}`

		req := httptest.NewRequest("POST", "/transactions", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.JSONEq(t, `{
			"error": "internal error"
		}`, string(body))

		mockTransactionService.AssertExpectations(t)
	})
}

func TestDeleteTransaction(t *testing.T) {
	t.Run("should delete transaction with status 204", func(t *testing.T) {
		app := fiber.New()

		mockTransactionService := new(services_mocks.TransactionServiceMock)
		mockTransactionService.On("DeleteTransaction", mock.Anything, uint(1)).Return(nil)

		controller := controllers.NewTransactionController(mockTransactionService)
		app.Delete("/transactions/:id", controller.DeleteTransaction)

		req := httptest.NewRequest("DELETE", "/transactions/1", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNoContent, resp.StatusCode)

		mockTransactionService.AssertExpectations(t)
	})

	t.Run("should return error with status 400 when id is invalid", func(t *testing.T) {
		app := fiber.New()

		mockTransactionService := new(services_mocks.TransactionServiceMock)

		controller := controllers.NewTransactionController(mockTransactionService)
		app.Delete("/transactions/:id", controller.DeleteTransaction)

		req := httptest.NewRequest("DELETE", "/transactions/invalid", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.JSONEq(t, `{
			"error": "`+utils_errors.IdInvalid+`"
		}`, string(body))

		mockTransactionService.AssertNotCalled(t, "DeleteTransaction")
	})

	t.Run("should return error with status 404 when transaction is not found", func(t *testing.T) {
		app := fiber.New()

		mockTransactionService := new(services_mocks.TransactionServiceMock)
		mockTransactionService.On("DeleteTransaction", mock.Anything, uint(1)).
			Return(utils_errors.ErrTransactionNotFound)

		controller := controllers.NewTransactionController(mockTransactionService)
		app.Delete("/transactions/:id", controller.DeleteTransaction)

		req := httptest.NewRequest("DELETE", "/transactions/1", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.JSONEq(t, `{
			"error": "`+utils_errors.TransactionNotFound+`"
		}`, string(body))

		mockTransactionService.AssertExpectations(t)
	})

	t.Run("should return error with status 500 when service fails", func(t *testing.T) {
		app := fiber.New()

		mockTransactionService := new(services_mocks.TransactionServiceMock)
		mockTransactionService.On("DeleteTransaction", mock.Anything, uint(1)).
			Return(errors.New("internal error"))

		controller := controllers.NewTransactionController(mockTransactionService)
		app.Delete("/transactions/:id", controller.DeleteTransaction)

		req := httptest.NewRequest("DELETE", "/transactions/1", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.JSONEq(t, `{
			"error": "internal error"
		}`, string(body))

		mockTransactionService.AssertExpectations(t)
	})
}

func TestUpdateTransaction(t *testing.T) {
	t.Run("should update transaction with status 200", func(t *testing.T) {
		app := fiber.New()

		mockTransactionService := new(services_mocks.TransactionServiceMock)
		mockTransactionService.On("UpdateTransaction", mock.Anything, uint(1), mock.Anything).Return(nil)

		controller := controllers.NewTransactionController(mockTransactionService)
		app.Put("/transactions/:id", controller.UpdateTransaction)

		requestBody := `{
			"title": "Updated Salary",
			"amount": 6000,
			"category": "Job",
			"type": "income"
		}`

		req := httptest.NewRequest("PUT", "/transactions/1", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.JSONEq(t, `{
			"message": "`+resolvers.TransactionUpdated+`"
		}`, string(body))

		mockTransactionService.AssertExpectations(t)
	})

	t.Run("should return error with status 400 when id is invalid", func(t *testing.T) {
		app := fiber.New()

		mockTransactionService := new(services_mocks.TransactionServiceMock)

		controller := controllers.NewTransactionController(mockTransactionService)
		app.Put("/transactions/:id", controller.UpdateTransaction)

		requestBody := `{
			"title": "Updated Salary",
			"amount": 6000,
			"category": "Job",
			"type": "income"
		}`

		req := httptest.NewRequest("PUT", "/transactions/invalid", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.JSONEq(t, `{
			"error": "`+utils_errors.IdInvalid+`"
		}`, string(body))

		mockTransactionService.AssertNotCalled(t, "UpdateTransaction")
	})

	t.Run("should return error with status 400 when title is empty", func(t *testing.T) {
		app := fiber.New()

		mockTransactionService := new(services_mocks.TransactionServiceMock)

		controller := controllers.NewTransactionController(mockTransactionService)
		app.Put("/transactions/:id", controller.UpdateTransaction)

		requestBody := `{
			"title": "",
			"amount": 6000,
			"category": "Job",
			"type": "income"
		}`

		req := httptest.NewRequest("PUT", "/transactions/1", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.JSONEq(t, `{
			"error": "`+utils_errors.TitleRequired+`"
		}`, string(body))

		mockTransactionService.AssertNotCalled(t, "UpdateTransaction")
	})

	t.Run("should return error with status 400 when amount is invalid", func(t *testing.T) {
		app := fiber.New()

		mockTransactionService := new(services_mocks.TransactionServiceMock)

		controller := controllers.NewTransactionController(mockTransactionService)
		app.Put("/transactions/:id", controller.UpdateTransaction)

		requestBody := `{
			"title": "Updated Salary",
			"amount": 0,
			"category": "Job",
			"type": "income"
		}`

		req := httptest.NewRequest("PUT", "/transactions/1", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.JSONEq(t, `{
			"error": "`+utils_errors.AmountRequired+`"
		}`, string(body))

		mockTransactionService.AssertNotCalled(t, "UpdateTransaction")
	})

	t.Run("should return error with status 400 when category is empty", func(t *testing.T) {
		app := fiber.New()

		mockTransactionService := new(services_mocks.TransactionServiceMock)

		controller := controllers.NewTransactionController(mockTransactionService)
		app.Put("/transactions/:id", controller.UpdateTransaction)

		requestBody := `{
			"title": "Updated Salary",
			"amount": 6000,
			"category": "",
			"type": "income"
		}`

		req := httptest.NewRequest("PUT", "/transactions/1", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.JSONEq(t, `{
			"error": "`+utils_errors.CategoryRequired+`"
		}`, string(body))

		mockTransactionService.AssertNotCalled(t, "UpdateTransaction")
	})

	t.Run("should return error with status 400 when type is invalid", func(t *testing.T) {
		app := fiber.New()

		mockTransactionService := new(services_mocks.TransactionServiceMock)

		controller := controllers.NewTransactionController(mockTransactionService)
		app.Put("/transactions/:id", controller.UpdateTransaction)

		requestBody := `{
			"title": "Updated Salary",
			"amount": 6000,
			"category": "Job",
			"type": "invalid"
		}`

		req := httptest.NewRequest("PUT", "/transactions/1", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.JSONEq(t, `{
			"error": "`+utils_errors.TypeInvalid+`"
		}`, string(body))

		mockTransactionService.AssertNotCalled(t, "UpdateTransaction")
	})

	t.Run("should return error with status 404 when transaction is not found", func(t *testing.T) {
		app := fiber.New()

		mockTransactionService := new(services_mocks.TransactionServiceMock)
		mockTransactionService.On("UpdateTransaction", mock.Anything, uint(1), mock.Anything).
			Return(utils_errors.ErrTransactionNotFound)

		controller := controllers.NewTransactionController(mockTransactionService)
		app.Put("/transactions/:id", controller.UpdateTransaction)

		requestBody := `{
			"title": "Updated Salary",
			"amount": 6000,
			"category": "Job",
			"type": "income"
		}`

		req := httptest.NewRequest("PUT", "/transactions/1", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.JSONEq(t, `{
			"error": "`+utils_errors.TransactionNotFound+`"
		}`, string(body))

		mockTransactionService.AssertExpectations(t)
	})

	t.Run("should return error with status 500 when service fails", func(t *testing.T) {
		app := fiber.New()

		mockTransactionService := new(services_mocks.TransactionServiceMock)
		mockTransactionService.On("UpdateTransaction", mock.Anything, uint(1), mock.Anything).
			Return(errors.New("internal error"))

		controller := controllers.NewTransactionController(mockTransactionService)
		app.Put("/transactions/:id", controller.UpdateTransaction)

		requestBody := `{
			"title": "Updated Salary",
			"amount": 6000,
			"category": "Job",
			"type": "income"
		}`

		req := httptest.NewRequest("PUT", "/transactions/1", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.JSONEq(t, `{
			"error": "internal error"
		}`, string(body))

		mockTransactionService.AssertExpectations(t)
	})
}
