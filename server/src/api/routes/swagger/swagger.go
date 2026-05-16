package routes

import (
	env "controle_financeiro/src/config/env"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func Init(app *fiber.App) {
	if env.AppEnv == "development" {
		app.Get("/swagger/*", fiberSwagger.WrapHandler)
	}
}
