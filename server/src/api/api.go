package api

import (
	_ "controle_financeiro/docs"
	health_route "controle_financeiro/src/api/routes/health"
	summary_route "controle_financeiro/src/api/routes/summary"
	swagger_route "controle_financeiro/src/api/routes/swagger"
	transaction_route "controle_financeiro/src/api/routes/transactions"
	env "controle_financeiro/src/config/env"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) error {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: env.CorsOrigin,
		AllowMethods: env.CorsMethod,
		AllowHeaders: env.CorsHeader,
	}))

	injectRoutes(app, db)

	port := fmt.Sprintf(":%s", env.Port)
	if err := app.Listen(port); err != nil {
		return fmt.Errorf("falha ao iniciar o servidor: %v", err)
	}

	return nil
}

func injectRoutes(app *fiber.App, db *gorm.DB) {
	health_route.Init(app)
	swagger_route.Init(app)

	transaction_route.Init(app, db)
	summary_route.Init(app, db)
}
