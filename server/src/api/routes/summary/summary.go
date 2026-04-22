package routes

import (
	summaryControllers "controle_financeiro/src/api/v1/controllers"
	sqliteConn "controle_financeiro/src/config/db/sqlite"
	sqlite "controle_financeiro/src/repositories/sqlite"
	"controle_financeiro/src/services"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	transactionRepository := sqlite.NewTransactionRepository(sqliteConn.DB)
	summaryService := services.NewSummaryService(transactionRepository)
	summaryController := summaryControllers.NewSummaryController(summaryService)

	v1 := app.Group("/v1")
	v1.Get("/summary", summaryController.GetSummary)
}
