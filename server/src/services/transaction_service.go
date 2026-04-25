package services

import (
	"context"
	"controle_financeiro/src/api/v1/dto"
	"controle_financeiro/src/models"
	repository_interfaces "controle_financeiro/src/repositories/interfaces"
	common "controle_financeiro/src/utils/common"
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

func (s *TransactionService) ListTransactions(ctx context.Context, filters dto.TransactionFilterDto) ([]dto.TransactionResponseDto, error) {
	transaction, err := s.SqliteTransactionRepositoryInterface.ListTransactions(ctx, filters)
	if err != nil {
		return nil, err
	}

	transactions := make([]dto.TransactionResponseDto, 0)

	for _, transaction := range transaction {
		transactions = append(transactions, dto.TransactionResponseDto{
			ID:          transaction.ID,
			Title:       transaction.Title,
			Description: transaction.Description,
			Amount:      transaction.Amount,
			Category:    transaction.Category,
			Type:        transaction.Type,
			CreatedAt:   transaction.CreatedAt,
		})
	}

	return transactions, nil
}

func (s *TransactionService) CreateTransaction(ctx context.Context, request dto.TransactionRequestDto) error {
	amount := request.Amount
	transactionType := common.TransactionTypeIncome

	if request.Type == common.TransactionTypeExpense {
		amount = -request.Amount
		transactionType = common.TransactionTypeExpense
	}

	transaction := models.Transaction{
		Title:       request.Title,
		Description: request.Description,
		Amount:      amount,
		Type:        transactionType,
		Category:    request.Category,
	}

	return s.SqliteTransactionRepositoryInterface.CreateTransaction(ctx, transaction)
}

func (s *TransactionService) DeleteTransaction(ctx context.Context, id uint) error {
	return s.SqliteTransactionRepositoryInterface.DeleteTransaction(ctx, id)
}

func (s *TransactionService) UpdateTransaction(ctx context.Context, id uint, request dto.TransactionRequestDto) error {
	amount := request.Amount
	transactionType := common.TransactionTypeIncome

	if request.Type == common.TransactionTypeExpense {
		amount = -request.Amount
		transactionType = common.TransactionTypeExpense
	}

	transaction := models.Transaction{
		Title:       request.Title,
		Description: request.Description,
		Amount:      amount,
		Type:        transactionType,
		Category:    request.Category,
	}

	return s.SqliteTransactionRepositoryInterface.UpdateTransaction(ctx, id, transaction)
}
