package main

import (
	"about/database"
	"about/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()
	routers.AboutRouters(app)
	app.Listen(":9092")
}
