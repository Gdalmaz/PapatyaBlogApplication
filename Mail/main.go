package main

import (
	"mail/database"
	"mail/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()
	routers.MailRouters(app)
	app.Listen(":9093")
}
