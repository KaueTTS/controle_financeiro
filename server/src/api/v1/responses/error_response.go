package responses

import (
	dto_shared "controle_financeiro/src/api/v1/dto/shared"
	shared_errors "controle_financeiro/src/shared/errors"

	"github.com/gofiber/fiber/v2"
)

func BadRequest(ctx *fiber.Ctx, message string, details []dto_shared.DetailErrorDto) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(dto_shared.ErrorDto{
		Message:     message,
		CodeMessage: shared_errors.BadRequest,
		Details:     details,
	})
}

func NotFound(ctx *fiber.Ctx, message string) error {
	return ctx.Status(fiber.StatusNotFound).JSON(dto_shared.ErrorDto{
		Message:     message,
		CodeMessage: shared_errors.NotFound,
	})
}

func InternalServerError(ctx *fiber.Ctx, message string) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(dto_shared.ErrorDto{
		Message:     message,
		CodeMessage: shared_errors.InternalServerError,
	})
}
