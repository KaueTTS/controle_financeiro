package routes

import (
	summaryControllers "controle_financeiro/src/api/v1/controllers"
	sqlite "controle_financeiro/src/repositories/sqlite"
	"controle_financeiro/src/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Init(app *fiber.App, db *gorm.DB) {
	summaryRepository := sqlite.NewSummaryRepository(db)
	summaryService := services.NewSummaryService(summaryRepository)
	summaryController := summaryControllers.NewSummaryController(summaryService)

	v1 := app.Group("/v1")
	v1.Get("/summary", summaryController.GetSummary)
}
