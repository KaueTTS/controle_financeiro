package health

import "github.com/gofiber/fiber/v2"

type Health struct {
	Status string `json:"status"`
}

//	HealthCheck godoc
//	@Summary		Verifica se a API está online
//	@Description	Retorna o status de funcionamento da API
//	@Tags			Health
//	@Produce		json
//	@Success		200	{object} Health
//	@Router			/health [get]
func healthRoute(c *fiber.Ctx) error {
	var health Health
	health.Status = "ok"

	return c.Status(fiber.StatusOK).JSON(health)
}

func Init(app *fiber.App) {
	app.Get("/health", healthRoute)
}
