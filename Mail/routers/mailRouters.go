package routers

import (
	"mail/controllers"

	"github.com/gofiber/fiber/v2"
)

func MailRouters(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	mail := v1.Group("/mail")

	mail.Post("/contact", controllers.SendMail)
}
