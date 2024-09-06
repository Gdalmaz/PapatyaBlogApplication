package main

import (
	"visitor/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routers.VisitorRouters(app)
	app.Listen(":9094")
}
