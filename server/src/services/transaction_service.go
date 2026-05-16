package services

import (
	"context"
	dto_shared "controle_financeiro/src/api/v1/dto/shared"
	dto_transaction "controle_financeiro/src/api/v1/dto/transaction"
	models "controle_financeiro/src/models"
	repository_interfaces "controle_financeiro/src/repositories/interfaces"
	"time"
)

type TransactionService struct {
	TransactionRepositoryInterface repository_interfaces.TransactionRepositoryInterface
}

func NewTransactionService(
	transactionRepository repository_interfaces.TransactionRepositoryInterface,
) *TransactionService {
	return &TransactionService{
		TransactionRepositoryInterface: transactionRepository,
	}
}

func (s *TransactionService) ListTransactions(ctx context.Context, filters dto_transaction.TransactionFilterDto) (dto_transaction.TransactionResponseDto, error) {
	if filters.Page <= 0 {
		filters.Page = 1
	}

	if filters.PerPage <= 0 {
		filters.PerPage = 10
	}

	if filters.PerPage > 100 {
		filters.PerPage = 100
	}

	transactionModel, total, err := s.TransactionRepositoryInterface.ListTransactions(ctx, filters)
	if err != nil {
		return dto_transaction.TransactionResponseDto{}, err
	}

	transactions := make([]dto_transaction.TransactionDto, 0, len(transactionModel))
	for _, transaction := range transactionModel {
		var deletedAt *time.Time
		if transaction.DeletedAt.Valid {
			deletedAt = &transaction.DeletedAt.Time
		}

		transactions = append(transactions, dto_transaction.TransactionDto{
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

	return dto_transaction.TransactionResponseDto{
		Pagination: dto_shared.PaginationDto{
			Page:      filters.Page,
			PerPage:   filters.PerPage,
			PageCount: pageCount,
			Total:     total,
		},
		Data: transactions,
	}, nil
}

func (s *TransactionService) CreateTransaction(ctx context.Context, request dto_transaction.TransactionRequestDto) error {
	transaction := models.Transaction{
		Title:       request.Title,
		Description: request.Description,
		Amount:      request.Amount,
		Type:        request.Type,
		Category:    request.Category,
	}

	return s.TransactionRepositoryInterface.CreateTransaction(ctx, transaction)
}

func (s *TransactionService) DeleteTransaction(ctx context.Context, id uint) error {
	return s.TransactionRepositoryInterface.DeleteTransaction(ctx, id)
}

func (s *TransactionService) UpdateTransaction(ctx context.Context, id uint, request dto_transaction.TransactionRequestDto) error {
	transaction := models.Transaction{
		Title:       request.Title,
		Description: request.Description,
		Amount:      request.Amount,
		Type:        request.Type,
		Category:    request.Category,
	}

	return s.TransactionRepositoryInterface.UpdateTransaction(ctx, id, transaction)
}
