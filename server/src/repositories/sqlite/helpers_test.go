package sqlite_repository_test

import (
	"controle_financeiro/src/models"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	require.NotNil(t, db)

	err = db.AutoMigrate(&models.Transaction{})
	require.NoError(t, err)

	return db
}
