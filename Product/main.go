package main

import (
	"product/database"
	"product/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()
	routers.ProductRouters(app)
	app.Listen(":9091")
}
