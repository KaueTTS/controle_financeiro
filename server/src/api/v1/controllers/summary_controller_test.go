package controllers_test

import (
	controllers "controle_financeiro/src/api/v1/controllers"
	dto_summary "controle_financeiro/src/api/v1/dto/summary"
	services_mocks "controle_financeiro/src/mocks/services"
	shared_errors "controle_financeiro/src/shared/errors"
	"errors"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetSummary(t *testing.T) {
	t.Run("should return summary with status 200", func(t *testing.T) {
		app := fiber.New()

		mockSummaryService := new(services_mocks.SummaryServiceMock)
		mockSummaryService.On("GetSummary", mock.Anything, dto_summary.SummaryFilterDto{}).Return(dto_summary.SummaryResponseDto{
			Income:  1000,
			Expense: 500,
			Balance: 500,
		}, nil)

		controller := controllers.NewSummaryController(mockSummaryService)
		app.Get("/summary", controller.GetSummary)

		req := httptest.NewRequest("GET", "/summary", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.JSONEq(t, `{
			"data": {
				"income": 1000,
				"expense": 500,
				"balance": 500
			}
		}`, string(body))

		mockSummaryService.AssertExpectations(t)
	})

	t.Run("should return error with status 500", func(t *testing.T) {
		app := fiber.New()

		mockSummaryService := new(services_mocks.SummaryServiceMock)
		mockSummaryService.On("GetSummary", mock.Anything, dto_summary.SummaryFilterDto{}).Return(dto_summary.SummaryResponseDto{}, errors.New("internal error"))

		controller := controllers.NewSummaryController(mockSummaryService)
		app.Get("/summary", controller.GetSummary)

		req := httptest.NewRequest("GET", "/summary", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.JSONEq(t, `{
			"message": "`+shared_errors.InternalServerErrorMessage+`",
			"codeMessage": "`+shared_errors.InternalServerError+`"
		}`, string(body))

		mockSummaryService.AssertExpectations(t)
	})
}
