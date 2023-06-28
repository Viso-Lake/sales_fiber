package routes

import (
	"sales_fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Auth routes
	app.Post("/cashiers/:cashierId/login", controllers.Login)
	app.Get("/cashiers/:cashierId/logout", controllers.Logout)
	app.Post("/cashiers/chashierId/passcode", controllers.Passcode)

	// Cashiers routes
	app.Post("/cashiers", controllers.CreateCashier)
	app.Get("/cashiers", controllers.CashiersList)
	app.Get("/cashiers/:cashierId", controllers.GetCashierDetails)
	app.Delete("/cashiers/:cashierId", controllers.DeleteCashier)
	app.Put("/cashiers/:cashierId", controllers.UpdateCashier)

}
