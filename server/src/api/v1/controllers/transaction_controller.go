package controllers

import (
	"controle_financeiro/src/api/v1/dto"
	servicesInterfaces "controle_financeiro/src/services/interfaces"
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
	var filters dto.TransactionFilterDto

	if err := ctx.QueryParser(&filters); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "query params inválidos",
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
			"error": "title é obrigatório",
		})
	}
	if request.Description == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "description é obrigatório",
		})
	}
	if request.Amount == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "amount não pode ser zero",
		})
	}
	if request.Category == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "category é obrigatório",
		})
	}
	if request.Type != "income" && request.Type != "expense" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "type deve ser income ou expense",
		})
	}

	err := c.transactionService.CreateTransaction(ctx.UserContext(), request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "transação criada com sucesso",
	})
}

func (c *TransactionController) DeleteTransaction(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id inválido",
		})
	}

	err = c.transactionService.DeleteTransaction(ctx.UserContext(), uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
