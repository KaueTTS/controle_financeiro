package main

import (
	api "controle_financeiro/src/api"
	sqlite_conn "controle_financeiro/src/config/db/sqlite"
	env "controle_financeiro/src/config/env"
	"fmt"

	"github.com/gofiber/fiber/v2/log"

	_ "controle_financeiro/docs"
)

// @title Controle Financeiro API
// @version 1.0
// @description API do sistema de controle financeiro

// @contact.name KauêTTS
// @contact.email kauebertaze2004@gmai.com

// @accept json
// @produce json

// @schemes http

// @host localhost:8080
// @BasePath /
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
