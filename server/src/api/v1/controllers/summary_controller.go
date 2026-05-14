package controllers

import (
	"controle_financeiro/src/api/v1/responses"
	servicesInterfaces "controle_financeiro/src/services/interfaces"
	utils_errors "controle_financeiro/src/utils/errors"

	"github.com/gofiber/fiber/v2"
)

type SummaryController struct {
	summaryService servicesInterfaces.SummaryServiceInterface
}

func NewSummaryController(summaryService servicesInterfaces.SummaryServiceInterface) *SummaryController {
	return &SummaryController{
		summaryService: summaryService,
	}
}

// GetSummary godoc
// @Summary Retorna o resumo financeiro
// @Description Cálculo automático de entradas, saídas e saldo total
// @Tags Summary
// @Success 200 {object} dto.SummaryResponseDto
// @failure 500 {object} dto.ErrorDto
// @Router /v1/summary [get]
func (c *SummaryController) GetSummary(ctx *fiber.Ctx) error {
	response, err := c.summaryService.GetSummary(ctx.UserContext())
	if err != nil {
		return responses.InternalServerError(ctx, utils_errors.InternalServerErrorMessage)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response,
	})
}
