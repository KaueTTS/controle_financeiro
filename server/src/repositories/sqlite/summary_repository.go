package sqlite_repository

import (
	"context"
	dto_summary "controle_financeiro/src/api/v1/dto/summary"
	models "controle_financeiro/src/models"
	shared_constants "controle_financeiro/src/shared/constants"

	"gorm.io/gorm"
)

type SummaryRepository struct {
	db *gorm.DB
}

func NewSummaryRepository(db *gorm.DB) *SummaryRepository {
	return &SummaryRepository{
		db: db,
	}
}

func (r *SummaryRepository) GetSummary(ctx context.Context, filters dto_summary.SummaryFilterDto) (dto_summary.SummaryResponseDto, error) {
	var response dto_summary.SummaryResponseDto

	query := r.db.WithContext(ctx).Model(&models.Transaction{})

	if filters.Search != "" {
		query = query.Where(
			"title LIKE ? OR description LIKE ?",
			"%"+filters.Search+"%",
			"%"+filters.Search+"%",
		)
	}

	if filters.Category != "" {
		query = query.Where("category = ?", filters.Category)
	}

	if filters.Type != "" {
		query = query.Where("type = ?", filters.Type)
	}

	if err := query.Select(
		`COALESCE(SUM(CASE WHEN type = ? THEN amount ELSE 0 END), 0) AS income,
		 COALESCE(SUM(CASE WHEN type = ? THEN amount ELSE 0 END), 0) AS expense`,
		shared_constants.TransactionTypeIncome,
		shared_constants.TransactionTypeExpense,
	).Scan(&response).Error; err != nil {
		return dto_summary.SummaryResponseDto{}, err
	}

	response.Balance = response.Income - response.Expense

	return response, nil
}
