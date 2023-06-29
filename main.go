package main

import (
	"fmt"
	db "sales_fiber/config"
	routes "sales_fiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Server started...")
	// db connection
	db.Connect()

	app := fiber.New()
	// app.Use(app)
	// Routing
	routes.Setup(app)
	app.Listen(":3000")

}
