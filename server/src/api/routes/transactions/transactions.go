package routes

import (
	controllers "controle_financeiro/src/api/v1/controllers"
	sqlite "controle_financeiro/src/repositories/sqlite"
	services "controle_financeiro/src/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Init(app *fiber.App, db *gorm.DB) {
	transactionRepository := sqlite.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepository)
	transactionController := controllers.NewTransactionController(transactionService)

	v1 := app.Group("/v1")
	v1.Get("/transactions", transactionController.ListTransactions)
	v1.Post("/transactions", transactionController.CreateTransaction)
	v1.Delete("/transactions/:id", transactionController.DeleteTransaction)
	v1.Put("/transactions/:id", transactionController.UpdateTransaction)
}
