package main

import (
	"otaqu/database"
	"otaqu/route"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.DatabaseInit()
	app := fiber.New()

	//initial route
	route.RouteInit(app)

	app.Listen(":8000")
}
