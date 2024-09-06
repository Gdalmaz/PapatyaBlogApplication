package routers

import (
	"visitor/controllers"

	"github.com/gofiber/fiber/v2"
)

func VisitorRouters(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	visitor := v1.Group("/visitor")
	visitor.Get("/get-all-products", controllers.GetAllPost)
	visitor.Get("get-all-extraurl", controllers.GetAllExtraImage)
}
