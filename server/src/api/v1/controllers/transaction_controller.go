package controllers

import (
	dto_shared "controle_financeiro/src/api/v1/dto/shared"
	dto_transaction "controle_financeiro/src/api/v1/dto/transaction"
	responses "controle_financeiro/src/api/v1/responses"
	validators "controle_financeiro/src/api/v1/validators"
	services_interfaces "controle_financeiro/src/services/interfaces"
	shared_constants "controle_financeiro/src/shared/constants"
	shared_errors "controle_financeiro/src/shared/errors"
	shared_http "controle_financeiro/src/shared/http"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TransactionController struct {
	transactionService services_interfaces.TransactionServiceInterface
}

func NewTransactionController(transactionService services_interfaces.TransactionServiceInterface) *TransactionController {
	return &TransactionController{
		transactionService: transactionService,
	}
}

// ListTransactions godoc
// @Summary Lista todas as transações
// @Description Visualização completa das movimentações financeiras cadastradas
// @Tags Transaction
// @Param search query string false "Buscar por título ou descrição "
// @Param type query string false "Tipo da transação" Enums(income, expense)
// @Param category query string false "Categoria da transação"
// @Param page query int false "Página atual" default(1)
// @Param perPage query int false "Quantidade de registros por página" default(10)
// @Success 200 {object} dto_transaction.TransactionResponseDto
// @Failure 400 {object} dto_shared.ErrorDto
// @Failure 500 {object} dto_shared.ErrorDto
// @Router /v1/transactions [get]
func (c *TransactionController) ListTransactions(ctx *fiber.Ctx) error {
	var filters dto_transaction.TransactionFilterDto
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

	response, err := c.transactionService.ListTransactions(ctx.UserContext(), filters)
	if err != nil {
		return responses.InternalServerError(ctx, shared_errors.InternalServerErrorMessage)
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

// CreateTransaction godoc
// @Summary Cadastrar transações
// @Description Cadastro de receitas e despesas
// @Tags Transaction
// @Param request body dto_transaction.TransactionRequestDto true "Dados da transação"
// @Success 201
// @Failure 400 {object} dto_shared.ErrorDto
// @Failure 500 {object} dto_shared.ErrorDto
// @Router /v1/transactions [post]
func (c *TransactionController) CreateTransaction(ctx *fiber.Ctx) error {
	var request dto_transaction.TransactionRequestDto
	if err := ctx.BodyParser(&request); err != nil {
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

	validationErrors := validators.ValidateTransactionRequest(request)
	if len(validationErrors) > 0 {
		return responses.BadRequest(
			ctx,
			shared_errors.MandatoryFieldMessage,
			validationErrors,
		)
	}

	err := c.transactionService.CreateTransaction(ctx.UserContext(), request)
	if err != nil {
		return responses.InternalServerError(ctx, shared_errors.InternalServerErrorMessage)
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": shared_http.TransactionCreated,
	})
}

// DeleteTransaction godoc
// @Summary Deletar transação
// @Description Remoção de transações cadastradas
// @Tags Transaction
// @Param id path int true "ID da transação"
// @Success 204 "No Content"
// @Failure 400 {object} dto_shared.ErrorDto
// @Failure 404 {object} dto_shared.ErrorDto
// @Failure 500 {object} dto_shared.ErrorDto
// @Router /v1/transactions/{id} [delete]
func (c *TransactionController) DeleteTransaction(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return responses.BadRequest(
			ctx,
			shared_errors.MandatoryFieldMessage,
			[]dto_shared.DetailErrorDto{
				{
					Field:   shared_constants.Id,
					Value:   shared_constants.Invalid,
					Message: shared_errors.IdInvalid,
				},
			},
		)
	}

	err = c.transactionService.DeleteTransaction(ctx.UserContext(), uint(id))
	if err != nil {
		if errors.Is(err, shared_errors.ErrTransactionNotFound) {
			return responses.NotFound(ctx, shared_errors.TransactionNotFoundMessage)
		}

		return responses.InternalServerError(ctx, shared_errors.InternalServerErrorMessage)
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}

// UpdateTransaction godoc
// @Summary Editar transação
// @Description Atualização de informações das transações
// @Tags Transaction
// @Param id path int true "ID da transação"
// @Param request body dto_transaction.TransactionRequestDto true "Dados atualizados da transação"
// @Success 200
// @Failure 400 {object} dto_shared.ErrorDto
// @Failure 404 {object} dto_shared.ErrorDto
// @Failure 500 {object} dto_shared.ErrorDto
// @Router /v1/transactions/{id} [put]
func (c *TransactionController) UpdateTransaction(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return responses.BadRequest(
			ctx,
			shared_errors.MandatoryFieldMessage,
			[]dto_shared.DetailErrorDto{
				{
					Field:   shared_constants.Id,
					Value:   shared_constants.Invalid,
					Message: shared_errors.IdInvalid,
				},
			},
		)
	}

	var request dto_transaction.TransactionRequestDto
	if err := ctx.BodyParser(&request); err != nil {
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

	validationErrors := validators.ValidateTransactionRequest(request)
	if len(validationErrors) > 0 {
		return responses.BadRequest(
			ctx,
			shared_errors.MandatoryFieldMessage,
			validationErrors,
		)
	}

	err = c.transactionService.UpdateTransaction(ctx.UserContext(), uint(id), request)
	if err != nil {
		if errors.Is(err, shared_errors.ErrTransactionNotFound) {
			return responses.NotFound(ctx, shared_errors.TransactionNotFoundMessage)
		}

		return responses.InternalServerError(ctx, shared_errors.InternalServerErrorMessage)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": shared_http.TransactionUpdated,
	})
}
