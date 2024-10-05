package main

import (
	"visitor/database"
	"visitor/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	
	app := fiber.New()
	database.Connect()
	routers.VisitorRouters(app)
	app.Listen(":9094")
}
