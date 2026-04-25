package sqlite_repository

import (
	"context"
	"controle_financeiro/src/api/v1/dto"
	"controle_financeiro/src/models"
	utils_errors "controle_financeiro/src/utils/errors"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (r *TransactionRepository) ListTransactions(ctx context.Context, filters dto.TransactionFilterDto) ([]models.Transaction, error) {
	var transactions []models.Transaction

	query := r.db.WithContext(ctx).Model(&models.Transaction{})

	if filters.Search != "" {
		query = query.Where("title LIKE ? OR description LIKE ?", "%"+filters.Search+"%", "%"+filters.Search+"%")
	}

	if filters.Category != "" {
		query = query.Where("category = ?", filters.Category)
	}

	if filters.Type == "income" {
		query = query.Where("amount > 0")
	}

	if filters.Type == "expense" {
		query = query.Where("amount < 0")
	}

	if err := query.Order("created_at desc").Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *TransactionRepository) CreateTransaction(ctx context.Context, transaction models.Transaction) error {
	return r.db.WithContext(ctx).Create(&transaction).Error
}

func (r *TransactionRepository) DeleteTransaction(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.Transaction{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils_errors.ErrTransactionNotFound
	}

	return nil
}

func (r *TransactionRepository) UpdateTransaction(ctx context.Context, id uint, transaction models.Transaction) error {
	result := r.db.WithContext(ctx).
		Model(&models.Transaction{}).
		Where("id = ?", id).
		Updates(transaction)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils_errors.ErrTransactionNotFound
	}

	return nil
}
