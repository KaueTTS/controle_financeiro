package sqlite_conn

import (
	env "controle_financeiro/src/config/env"
	"fmt"
	"os"
	"path/filepath"

	"controle_financeiro/src/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	dir := filepath.Dir(env.DatabaseURL)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("criar diretório do banco de dados: %w", err)
	}

	db, err := gorm.Open(sqlite.Open(env.DatabaseURL), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("abrir banco de dados: %w", err)
	}

	if err := db.AutoMigrate(&models.Transaction{}); err != nil {
		return fmt.Errorf("migrar banco de dados: %w", err)
	}

	DB = db

	return nil
}
