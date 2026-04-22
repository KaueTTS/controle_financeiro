package health

import "github.com/gofiber/fiber/v2"

type Health struct {
	Status string `json:"status"`
}

func healthRoute(c *fiber.Ctx) error {
	var health Health
	health.Status = "ok"

	return c.Status(fiber.StatusOK).JSON(health)
}

func Init(app *fiber.App) {
	app.Get("/health", healthRoute)
}
