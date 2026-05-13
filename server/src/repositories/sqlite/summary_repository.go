package sqlite_repository

import (
	"context"
	"controle_financeiro/src/api/v1/dto"
	"controle_financeiro/src/models"
	"controle_financeiro/src/utils/common"

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

func (r *SummaryRepository) GetSummary(ctx context.Context) (dto.SummaryResponseDto, error) {
	var income float64
	var expense float64

	if err := r.db.WithContext(ctx).
		Model(&models.Transaction{}).
		Where("type = ?", common.TransactionTypeIncome).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&income).Error; err != nil {
		return dto.SummaryResponseDto{}, err
	}

	if err := r.db.WithContext(ctx).
		Model(&models.Transaction{}).
		Where("type = ?", common.TransactionTypeExpense).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&expense).Error; err != nil {
		return dto.SummaryResponseDto{}, err
	}

	return dto.SummaryResponseDto{
		Income:  income,
		Expense: expense,
		Balance: income - expense,
	}, nil
}
