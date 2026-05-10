package controllers

import (
	"controle_financeiro/src/api/v1/dto"
	servicesInterfaces "controle_financeiro/src/services/interfaces"
	"controle_financeiro/src/utils/common"
	utils_errors "controle_financeiro/src/utils/errors"
	resolvers "controle_financeiro/src/utils/resolvers"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TransactionController struct {
	transactionService servicesInterfaces.TransactionServiceInterface
}

func NewTransactionController(transactionService servicesInterfaces.TransactionServiceInterface) *TransactionController {
	return &TransactionController{
		transactionService: transactionService,
	}
}

func (c *TransactionController) ListTransactions(ctx *fiber.Ctx) error {
	var filters dto.FilterDto
	if err := ctx.QueryParser(&filters); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response, err := c.transactionService.ListTransactions(ctx.UserContext(), filters)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response,
	})
}

func (c *TransactionController) CreateTransaction(ctx *fiber.Ctx) error {
	var request dto.TransactionRequestDto

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if request.Title == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils_errors.TitleRequired,
		})
	}
	if request.Amount <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils_errors.AmountRequired,
		})
	}
	if request.Category == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils_errors.CategoryRequired,
		})
	}
	if request.Type != common.TransactionTypeIncome && request.Type != common.TransactionTypeExpense {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils_errors.TypeInvalid,
		})
	}

	err := c.transactionService.CreateTransaction(ctx.UserContext(), request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": resolvers.TransactionCreated,
	})
}

func (c *TransactionController) DeleteTransaction(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils_errors.IdInvalid,
		})
	}

	err = c.transactionService.DeleteTransaction(ctx.UserContext(), uint(id))
	if err != nil {
		if errors.Is(err, utils_errors.ErrTransactionNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": utils_errors.TransactionNotFound,
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *TransactionController) UpdateTransaction(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils_errors.IdInvalid,
		})
	}

	var request dto.TransactionRequestDto
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if request.Title == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils_errors.TitleRequired,
		})
	}
	if request.Amount <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils_errors.AmountRequired,
		})
	}
	if request.Category == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils_errors.CategoryRequired,
		})
	}
	if request.Type != common.TransactionTypeIncome && request.Type != common.TransactionTypeExpense {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils_errors.TypeInvalid,
		})
	}

	err = c.transactionService.UpdateTransaction(ctx.UserContext(), uint(id), request)
	if err != nil {
		if errors.Is(err, utils_errors.ErrTransactionNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": utils_errors.TransactionNotFound,
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": resolvers.TransactionUpdated,
	})
}
