package sqlite_conn

import (
	env "controle_financeiro/src/config/env"
	"fmt"
	"os"
	"path/filepath"

	"controle_financeiro/src/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
	dir := filepath.Dir(env.DatabaseURL)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, fmt.Errorf("erro ao criar diretório do banco de dados: %w", err)
	}

	db, err := gorm.Open(sqlite.Open(env.DatabaseURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar no sqlite: %w", err)
	}

	if err := db.AutoMigrate(
		&models.Transaction{},
	); err != nil {
		return nil, fmt.Errorf("erro ao migrar banco de dados: %w", err)
	}

	return db, nil
}
