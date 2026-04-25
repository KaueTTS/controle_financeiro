package routes

import (
	"controle_financeiro/src/api/v1/controllers"
	sqliteConn "controle_financeiro/src/config/db/sqlite"
	sqlite "controle_financeiro/src/repositories/sqlite"
	"controle_financeiro/src/services"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	transactionRepository := sqlite.NewTransactionRepository(sqliteConn.DB)

	transactionService := services.NewTransactionService(
		transactionRepository,
	)

	transactionController := controllers.NewTransactionController(transactionService)

	v1 := app.Group("/v1")
	v1.Get("/transactions", transactionController.ListTransactions)
	v1.Post("/transactions", transactionController.CreateTransaction)
	v1.Delete("/transactions/:id", transactionController.DeleteTransaction)
	v1.Put("/transactions/:id", transactionController.UpdateTransaction)
}
