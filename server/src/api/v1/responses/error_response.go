package responses

import (
	"controle_financeiro/src/api/v1/dto"
	utils_errors "controle_financeiro/src/utils/errors"

	"github.com/gofiber/fiber/v2"
)

func BadRequest(ctx *fiber.Ctx, message string, details []dto.DetailErrorDto) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorDto{
		Message:     message,
		CodeMessage: utils_errors.BadRequest,
		Details:     details,
	})
}

func NotFound(ctx *fiber.Ctx, message string) error {
	return ctx.Status(fiber.StatusNotFound).JSON(dto.ErrorDto{
		Message:     message,
		CodeMessage: utils_errors.NotFound,
	})
}

func InternalServerError(ctx *fiber.Ctx, message string) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorDto{
		Message:     message,
		CodeMessage: utils_errors.InternalServerError,
	})
}
