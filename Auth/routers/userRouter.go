package routers

import (
	"auth/controllers"
	"auth/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	user := v1.Group("/user")
	user.Post("/sign-up", controllers.SignUp)
	user.Post("/sign-in", controllers.SignIn)
	user.Put("update-pass", middleware.TokenControl(), controllers.UpdatePassword)
	user.Put("/update-acc", middleware.TokenControl(), controllers.UpdateAccount)
	user.Delete("/delete-acc", middleware.TokenControl(), controllers.DeleteAccount)
}
