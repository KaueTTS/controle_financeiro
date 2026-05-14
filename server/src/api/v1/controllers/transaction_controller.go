package controllers

import (
	"controle_financeiro/src/api/v1/dto"
	"controle_financeiro/src/api/v1/responses"
	"controle_financeiro/src/api/v1/validators"
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

// ListTransactions godoc
// @Summary Lista todas as transações
// @Description Visualização completa das movimentações financeiras cadastradas
// @Tags Transaction
// @Success 200 {object} dto.TransactionResponseDto
// @failure 400 {object} dto.ErrorDto
// @failure 500 {object} dto.ErrorDto
// @Router /v1/transactions [get]
func (c *TransactionController) ListTransactions(ctx *fiber.Ctx) error {
	var filters dto.FilterDto
	if err := ctx.QueryParser(&filters); err != nil {
		return responses.BadRequest(
			ctx,
			utils_errors.InvalidRequestMessage,
			[]dto.DetailErrorDto{
				{
					Field:   "",
					Value:   "",
					Message: err.Error(),
				},
			},
		)
	}

	response, err := c.transactionService.ListTransactions(ctx.UserContext(), filters)
	if err != nil {
		return responses.InternalServerError(ctx, utils_errors.InternalServerErrorMessage)
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

// CreateTransaction godoc
// @Summary
// @Description
// @Tags Transaction
// @failure 400 {object} dto.ErrorDto
// @failure 500 {object} dto.ErrorDto
// @Router /v1/transactions [post]
func (c *TransactionController) CreateTransaction(ctx *fiber.Ctx) error {
	var request dto.TransactionRequestDto
	if err := ctx.BodyParser(&request); err != nil {
		return responses.BadRequest(
			ctx,
			utils_errors.InvalidRequestMessage,
			[]dto.DetailErrorDto{
				{
					Field:   "",
					Value:   "",
					Message: err.Error(),
				},
			},
		)
	}

	validationErrors := validators.ValidateTransactionRequest(request)
	if len(validationErrors) > 0 {
		return responses.BadRequest(
			ctx,
			utils_errors.MandatoryFieldMessage,
			validationErrors,
		)
	}

	err := c.transactionService.CreateTransaction(ctx.UserContext(), request)
	if err != nil {
		return responses.InternalServerError(ctx, utils_errors.InternalServerErrorMessage)
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": resolvers.TransactionCreated,
	})
}

// DeleteTransaction godoc
// @Summary
// @Description
// @Tags Transaction
// @failure 400 {object} dto.ErrorDto
// @failure 404 {object} dto.ErrorDto
// @failure 500 {object} dto.ErrorDto
// @Router /v1/transactions/:id [delete]
func (c *TransactionController) DeleteTransaction(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return responses.BadRequest(
			ctx,
			utils_errors.MandatoryFieldMessage,
			[]dto.DetailErrorDto{
				{
					Field:   common.Id,
					Value:   common.Invalid,
					Message: utils_errors.IdInvalid,
				},
			},
		)
	}

	err = c.transactionService.DeleteTransaction(ctx.UserContext(), uint(id))
	if err != nil {
		if errors.Is(err, utils_errors.ErrTransactionNotFound) {
			return responses.NotFound(ctx, utils_errors.TransactionNotFoundMessage)
		}

		return responses.InternalServerError(ctx, utils_errors.InternalServerErrorMessage)
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}

// UpdateTransaction godoc
// @Summary
// @Description
// @Tags Transaction
// @failure 400 {object} dto.ErrorDto
// @failure 404 {object} dto.ErrorDto
// @failure 500 {object} dto.ErrorDto
// @Router /v1/transactions/:id [put]
func (c *TransactionController) UpdateTransaction(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return responses.BadRequest(
			ctx,
			utils_errors.MandatoryFieldMessage,
			[]dto.DetailErrorDto{
				{
					Field:   common.Id,
					Value:   common.Invalid,
					Message: utils_errors.IdInvalid,
				},
			},
		)
	}

	var request dto.TransactionRequestDto
	if err := ctx.BodyParser(&request); err != nil {
		return responses.BadRequest(
			ctx,
			utils_errors.InvalidRequestMessage,
			[]dto.DetailErrorDto{
				{
					Field:   "",
					Value:   "",
					Message: err.Error(),
				},
			},
		)
	}

	validationErrors := validators.ValidateTransactionRequest(request)
	if len(validationErrors) > 0 {
		return responses.BadRequest(
			ctx,
			utils_errors.MandatoryFieldMessage,
			validationErrors,
		)
	}

	err = c.transactionService.UpdateTransaction(ctx.UserContext(), uint(id), request)
	if err != nil {
		if errors.Is(err, utils_errors.ErrTransactionNotFound) {
			return responses.NotFound(ctx, utils_errors.TransactionNotFoundMessage)
		}

		return responses.InternalServerError(ctx, utils_errors.InternalServerErrorMessage)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": resolvers.TransactionUpdated,
	})
}
