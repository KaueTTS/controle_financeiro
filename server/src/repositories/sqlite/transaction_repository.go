package sqlite_repository

import (
	"context"
	"controle_financeiro/src/api/v1/dto"
	"controle_financeiro/src/models"
	utils_errors "controle_financeiro/src/utils/errors"
	"time"

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

func (r *TransactionRepository) ListTransactions(ctx context.Context, filters dto.FilterDto) ([]models.Transaction, int64, error) {
	var transactions []models.Transaction
	var total int64

	query := r.db.WithContext(ctx).Model(&models.Transaction{})

	if filters.Search != "" {
		query = query.Where(
			"title LIKE ? OR description LIKE ?",
			"%"+filters.Search+"%",
			"%"+filters.Search+"%")
	}

	if filters.Category != "" {
		query = query.Where("category = ?", filters.Category)
	}

	if filters.Type != "" {
		query = query.Where("type = ?", filters.Type)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (filters.Page - 1) * filters.PerPage

	if err := query.
		Order("created_at desc").
		Limit(filters.PerPage).
		Offset(offset).
		Find(&transactions).Error; err != nil {
		return nil, 0, err
	}

	return transactions, total, nil
}

func (r *TransactionRepository) CreateTransaction(ctx context.Context, transaction models.Transaction) error {
	return r.db.WithContext(ctx).Create(&transaction).Error
}

func (r *TransactionRepository) DeleteTransaction(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).
		Model(&models.Transaction{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Update("deleted_at", time.Now())

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
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(map[string]interface{}{
			"title":       transaction.Title,
			"description": transaction.Description,
			"amount":      transaction.Amount,
			"type":        transaction.Type,
			"category":    transaction.Category,
			"updated_at":  time.Now(),
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils_errors.ErrTransactionNotFound
	}

	return nil
}
