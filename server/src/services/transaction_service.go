package services

import (
	"context"
	"controle_financeiro/src/api/v1/dto"
	"controle_financeiro/src/models"
	repository_interfaces "controle_financeiro/src/repositories/sqlite/interfaces"
	"time"
)

type TransactionService struct {
	SqliteTransactionRepositoryInterface repository_interfaces.SqliteTransactionRepositoryInterface
}

func NewTransactionService(
	sqliteTransactionRepository repository_interfaces.SqliteTransactionRepositoryInterface,
) *TransactionService {
	return &TransactionService{
		SqliteTransactionRepositoryInterface: sqliteTransactionRepository,
	}
}

func (s *TransactionService) ListTransactions(ctx context.Context, filters dto.FilterDto) (dto.TransactionResponseDto, error) {
	if filters.Page <= 0 {
		filters.Page = 1
	}

	if filters.PerPage <= 0 {
		filters.PerPage = 10
	}

	if filters.PerPage > 100 {
		filters.PerPage = 100
	}

	transactionModel, total, err := s.SqliteTransactionRepositoryInterface.ListTransactions(ctx, filters)
	if err != nil {
		return dto.TransactionResponseDto{}, err
	}

	transactions := make([]dto.TransactionDto, 0, len(transactionModel))
	for _, transaction := range transactionModel {
		var deletedAt *time.Time
		if transaction.DeletedAt.Valid {
			deletedAt = &transaction.DeletedAt.Time
		}

		transactions = append(transactions, dto.TransactionDto{
			ID:          transaction.ID,
			Title:       transaction.Title,
			Description: transaction.Description,
			Amount:      transaction.Amount,
			Category:    transaction.Category,
			Type:        transaction.Type,
			CreatedAt:   transaction.CreatedAt,
			UpdatedAt:   &transaction.UpdatedAt,
			DeletedAt:   deletedAt,
		})
	}

	pageCount := int(total) / filters.PerPage
	if int(total)%filters.PerPage != 0 {
		pageCount++
	}

	return dto.TransactionResponseDto{
		Pagination: dto.PaginationDto{
			Page:      filters.Page,
			PerPage:   filters.PerPage,
			PageCount: pageCount,
			Total:     total,
		},
		Data: transactions,
	}, nil
}

func (s *TransactionService) CreateTransaction(ctx context.Context, request dto.TransactionRequestDto) error {
	transaction := models.Transaction{
		Title:       request.Title,
		Description: request.Description,
		Amount:      request.Amount,
		Type:        request.Type,
		Category:    request.Category,
	}

	return s.SqliteTransactionRepositoryInterface.CreateTransaction(ctx, transaction)
}

func (s *TransactionService) DeleteTransaction(ctx context.Context, id uint) error {
	return s.SqliteTransactionRepositoryInterface.DeleteTransaction(ctx, id)
}

func (s *TransactionService) UpdateTransaction(ctx context.Context, id uint, request dto.TransactionRequestDto) error {
	transaction := models.Transaction{
		Title:       request.Title,
		Description: request.Description,
		Amount:      request.Amount,
		Type:        request.Type,
		Category:    request.Category,
	}

	return s.SqliteTransactionRepositoryInterface.UpdateTransaction(ctx, id, transaction)
}
