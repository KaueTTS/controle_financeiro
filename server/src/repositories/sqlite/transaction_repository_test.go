package sqlite_repository_test

import (
	"context"
	dto_transaction "controle_financeiro/src/api/v1/dto/transaction"
	models "controle_financeiro/src/models"
	sqlite_repository "controle_financeiro/src/repositories/sqlite"
	shared_errors "controle_financeiro/src/shared/errors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransactionRepository(t *testing.T) {
	db := setupTestDB(t)
	repository := sqlite_repository.NewTransactionRepository(db)

	transaction := models.Transaction{
		Title:    "Salário",
		Amount:   5000,
		Type:     "income",
		Category: "Trabalho",
	}

	err := repository.CreateTransaction(context.Background(), transaction)

	assert.NoError(t, err)

	var count int64
	db.Model(&models.Transaction{}).Count(&count)

	assert.Equal(t, int64(1), count)
}

func TestListTransactionsRepository(t *testing.T) {
	db := setupTestDB(t)
	repository := sqlite_repository.NewTransactionRepository(db)

	description := "Compra do mês"

	db.Create(&models.Transaction{
		Title:       "Mercado",
		Description: &description,
		Amount:      300,
		Type:        "expense",
		Category:    "Alimentação",
	})

	db.Create(&models.Transaction{
		Title:    "Salário",
		Amount:   5000,
		Type:     "income",
		Category: "Trabalho",
	})

	filters := dto_transaction.TransactionFilterDto{
		Search:   "Mercado",
		Type:     "expense",
		Category: "Alimentação",
		Page:     1,
		PerPage:  10,
	}

	transactions, total, err := repository.ListTransactions(context.Background(), filters)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), total)
	assert.Len(t, transactions, 1)
	assert.Equal(t, "Mercado", transactions[0].Title)
	assert.Equal(t, "expense", transactions[0].Type)
	assert.Equal(t, "Alimentação", transactions[0].Category)
}

func TestDeleteTransactionRepository(t *testing.T) {
	t.Run("should delete transaction", func(t *testing.T) {
		db := setupTestDB(t)
		repository := sqlite_repository.NewTransactionRepository(db)

		transaction := models.Transaction{
			Title:    "Mercado",
			Amount:   300,
			Type:     "expense",
			Category: "Alimentação",
		}

		db.Create(&transaction)

		err := repository.DeleteTransaction(context.Background(), transaction.ID)

		assert.NoError(t, err)

		var count int64
		db.Model(&models.Transaction{}).Where("id = ?", transaction.ID).Count(&count)

		assert.Equal(t, int64(0), count)
	})

	t.Run("should return error when transaction does not exist", func(t *testing.T) {
		db := setupTestDB(t)
		repository := sqlite_repository.NewTransactionRepository(db)

		err := repository.DeleteTransaction(context.Background(), 999)

		assert.Error(t, err)
		assert.True(t, errors.Is(err, shared_errors.ErrTransactionNotFound))
	})
}

func TestUpdateTransactionRepository(t *testing.T) {
	db := setupTestDB(t)
	repository := sqlite_repository.NewTransactionRepository(db)

	transaction := models.Transaction{
		Title:    "Mercado",
		Amount:   300,
		Type:     "expense",
		Category: "Alimentação",
	}

	db.Create(&transaction)

	update := models.Transaction{
		Title:    "Mercado atualizado",
		Amount:   350,
		Type:     "expense",
		Category: "Casa",
	}

	err := repository.UpdateTransaction(context.Background(), transaction.ID, update)

	assert.NoError(t, err)

	var result models.Transaction
	db.First(&result, transaction.ID)

	assert.Equal(t, "Mercado atualizado", result.Title)
	assert.Equal(t, float64(350), result.Amount)
	assert.Equal(t, "Casa", result.Category)
}
