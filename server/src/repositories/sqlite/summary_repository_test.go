package sqlite_repository_test

import (
	"context"
	dto_summary "controle_financeiro/src/api/v1/dto/summary"
	models "controle_financeiro/src/models"
	sqlite_repository "controle_financeiro/src/repositories/sqlite"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSummaryRepository(t *testing.T) {
	db := setupTestDB(t)
	repository := sqlite_repository.NewSummaryRepository(db)

	db.Create(&models.Transaction{
		Title:    "Salário",
		Amount:   5000,
		Type:     "income",
		Category: "Trabalho",
	})

	db.Create(&models.Transaction{
		Title:    "Freela",
		Amount:   1000,
		Type:     "income",
		Category: "Extra",
	})

	db.Create(&models.Transaction{
		Title:    "Mercado",
		Amount:   700,
		Type:     "expense",
		Category: "Alimentação",
	})

	response, err := repository.GetSummary(context.Background(), dto_summary.SummaryFilterDto{})

	assert.NoError(t, err)
	assert.Equal(t, float64(6000), response.Income)
	assert.Equal(t, float64(700), response.Expense)
	assert.Equal(t, float64(5300), response.Balance)
}
