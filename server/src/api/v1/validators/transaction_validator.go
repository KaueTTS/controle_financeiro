package validators

import (
	dto_shared "controle_financeiro/src/api/v1/dto/shared"
	dto_transaction "controle_financeiro/src/api/v1/dto/transaction"
	shared_constants "controle_financeiro/src/shared/constants"
	shared_errors "controle_financeiro/src/shared/errors"
	"strings"
)

func ValidateTransactionRequest(request dto_transaction.TransactionRequestDto) []dto_shared.DetailErrorDto {
	var errors []dto_shared.DetailErrorDto

	if strings.TrimSpace(request.Title) == "" {
		errors = append(errors, dto_shared.DetailErrorDto{
			Field:   shared_constants.Title,
			Value:   shared_constants.Mandatory,
			Message: shared_errors.TitleRequired,
		})
	}

	if request.Amount <= 0 {
		errors = append(errors, dto_shared.DetailErrorDto{
			Field:   shared_constants.Amount,
			Value:   shared_constants.GreaterThanZero,
			Message: shared_errors.AmountRequired,
		})
	}

	if strings.TrimSpace(request.Category) == "" {
		errors = append(errors, dto_shared.DetailErrorDto{
			Field:   shared_constants.Category,
			Value:   shared_constants.Mandatory,
			Message: shared_errors.CategoryRequired,
		})
	}

	if request.Type != shared_constants.TransactionTypeIncome && request.Type != shared_constants.TransactionTypeExpense {
		errors = append(errors, dto_shared.DetailErrorDto{
			Field:   shared_constants.Type,
			Value:   shared_constants.Invalid,
			Message: shared_errors.TypeInvalid,
		})
	}

	return errors
}
