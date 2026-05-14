package api

import (
	healthRoute "controle_financeiro/src/api/routes/health"
	summaryRoute "controle_financeiro/src/api/routes/summary"
	transactionRoute "controle_financeiro/src/api/routes/transactions"
	env "controle_financeiro/src/config/env"
	"fmt"

	_ "controle_financeiro/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) error {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: env.FrontendCorsOrigin,
		AllowMethods: "*",
		AllowHeaders: "*",
	}))

	if env.AppEnv != "prod" {
		app.Get("/swagger/*", fiberSwagger.WrapHandler)
	}

	injectRoutes(app, db)

	port := fmt.Sprintf(":%s", env.Port)
	if err := app.Listen(port); err != nil {
		return fmt.Errorf("falha ao iniciar o servidor: %v", err)
	}

	return nil
}

func injectRoutes(app *fiber.App, db *gorm.DB) {
	healthRoute.Init(app)

	transactionRoute.Init(app, db)
	summaryRoute.Init(app, db)
}
