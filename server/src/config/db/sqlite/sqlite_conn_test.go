package sqlite_conn_test

import (
	sqlite_conn "controle_financeiro/src/config/db/sqlite"
	env "controle_financeiro/src/config/env"
	"os"
	"path/filepath"
	"testing"
)

func TestInit(t *testing.T) {
	t.Run("should be successfully initialized", func(t *testing.T) {
		tempDir := t.TempDir()

		oldDatabaseURL := env.DatabaseURL
		env.DatabaseURL = filepath.Join(tempDir, "test.db")

		t.Cleanup(func() {
			env.DatabaseURL = oldDatabaseURL
		})

		db, err := sqlite_conn.Init()

		if err != nil {
			t.Fatalf("erro inesperado ao inicializar banco: %v", err)
		}

		if db == nil {
			t.Fatal("esperava que db não fosse nil")
		}

		sqlDB, err := db.DB()
		if err != nil {
			t.Fatalf("erro ao recuperar conexão sql: %v", err)
		}

		t.Cleanup(func() {
			_ = sqlDB.Close()
		})

		if _, err := os.Stat(env.DatabaseURL); err != nil {
			t.Fatalf("esperava que o arquivo do banco fosse criado: %v", err)
		}

		if !db.Migrator().HasTable("transactions") {
			t.Fatalf("esperava que a tabela transactions existisse")
		}
	})

	t.Run("should return error when database directory cannot be created", func(t *testing.T) {
		tempDir := t.TempDir()

		filePath := filepath.Join(tempDir, "not-a-dir")

		if err := os.WriteFile(filePath, []byte("content"), 0o644); err != nil {
			t.Fatalf("erro ao criar arquivo temporário: %v", err)
		}

		oldDatabaseURL := env.DatabaseURL
		env.DatabaseURL = filepath.Join(filePath, "test.db")

		t.Cleanup(func() {
			env.DatabaseURL = oldDatabaseURL
		})

		db, err := sqlite_conn.Init()

		if err == nil {
			t.Fatal("esperava erro ao inicializar banco")
		}

		if db != nil {
			t.Fatal("esperava que db fosse nil quando ocorresse erro")
		}
	})
}
