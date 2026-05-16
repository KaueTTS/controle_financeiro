package controllers

import (
	dto_shared "controle_financeiro/src/api/v1/dto/shared"
	dto_summary "controle_financeiro/src/api/v1/dto/summary"
	responses "controle_financeiro/src/api/v1/responses"
	services_interfaces "controle_financeiro/src/services/interfaces"
	shared_errors "controle_financeiro/src/shared/errors"

	"github.com/gofiber/fiber/v2"
)

type SummaryController struct {
	summaryService services_interfaces.SummaryServiceInterface
}

func NewSummaryController(summaryService services_interfaces.SummaryServiceInterface) *SummaryController {
	return &SummaryController{
		summaryService: summaryService,
	}
}

// GetSummary godoc
// @Summary Retorna o resumo financeiro
// @Description Cálculo automático de entradas, saídas e saldo total
// @Tags Summary
// @Param search query string false "Buscar por título ou descrição"
// @Param type query string false "Tipo da transação" Enums(income, expense)
// @Param category query string false "Categoria da transação"
// @Success 200 {object} dto_summary.SummaryResponseDto
// @Failure 400 {object} dto_shared.ErrorDto
// @Failure 500 {object} dto_shared.ErrorDto
// @Router /v1/summary [get]
func (c *SummaryController) GetSummary(ctx *fiber.Ctx) error {
	var filters dto_summary.SummaryFilterDto
	if err := ctx.QueryParser(&filters); err != nil {
		return responses.BadRequest(
			ctx,
			shared_errors.InvalidRequestMessage,
			[]dto_shared.DetailErrorDto{
				{
					Field:   "",
					Value:   "",
					Message: err.Error(),
				},
			},
		)
	}

	response, err := c.summaryService.GetSummary(ctx.UserContext(), filters)
	if err != nil {
		return responses.InternalServerError(ctx, shared_errors.InternalServerErrorMessage)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response,
	})
}
