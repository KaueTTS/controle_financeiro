package validators

import (
	"controle_financeiro/src/api/v1/dto"
	"controle_financeiro/src/utils/common"
	utils_errors "controle_financeiro/src/utils/errors"
	"strings"
)

func ValidateTransactionRequest(request dto.TransactionRequestDto) []dto.DetailErrorDto {
	var errors []dto.DetailErrorDto

	if strings.TrimSpace(request.Title) == "" {
		errors = append(errors, dto.DetailErrorDto{
			Field:   common.Title,
			Value:   common.Mandatory,
			Message: utils_errors.TitleRequired,
		})
	}

	if request.Amount <= 0 {
		errors = append(errors, dto.DetailErrorDto{
			Field:   common.Amount,
			Value:   common.GreaterThanZero,
			Message: utils_errors.AmountRequired,
		})
	}

	if strings.TrimSpace(request.Category) == "" {
		errors = append(errors, dto.DetailErrorDto{
			Field:   common.Category,
			Value:   common.Mandatory,
			Message: utils_errors.CategoryRequired,
		})
	}

	if request.Type != common.TransactionTypeIncome && request.Type != common.TransactionTypeExpense {
		errors = append(errors, dto.DetailErrorDto{
			Field:   common.Type,
			Value:   common.Invalid,
			Message: utils_errors.TypeInvalid,
		})
	}

	return errors
}
