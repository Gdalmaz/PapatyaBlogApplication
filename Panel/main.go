package main

import (
	"panel/database"
	"panel/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()
	routers.PanelRouters(app)
	app.Listen(":9093")
}