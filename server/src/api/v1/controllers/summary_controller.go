package controllers

import (
	servicesInterfaces "controle_financeiro/src/services/interfaces"

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

func (c *SummaryController) GetSummary(ctx *fiber.Ctx) error {
	response, err := c.summaryService.GetSummary(ctx.UserContext())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response,
	})
}
