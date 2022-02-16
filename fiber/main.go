package main

import (
	"fiber/configs"
	"fiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(app)

	app.Listen(":6000")
}
