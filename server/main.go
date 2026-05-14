package main

import (
	api "controle_financeiro/src/api"
	sqlite_conn "controle_financeiro/src/config/db/sqlite"
	env "controle_financeiro/src/config/env"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("falha ao iniciar aplicação: %v", err)
	}
}

func run() error {
	if err := env.Init(); err != nil {
		return fmt.Errorf("erro ao inicializar variáveis de ambiente: %w", err)
	}

	db, err := sqlite_conn.Init()
	if err != nil {
		return fmt.Errorf("erro ao inicializar sqlite: %w", err)
	}

	if err := api.Init(db); err != nil {
		return fmt.Errorf("erro ao iniciar api: %w", err)
	}

	return nil
}
