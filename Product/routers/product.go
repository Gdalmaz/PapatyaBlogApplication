package routers

import (
	"product/controllers"
	"product/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductRouters(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	product := v1.Group("/product")

	product.Post("/add-post", middleware.TokenControl(), controllers.AddProduct)
	product.Put("/update-product/:id", middleware.TokenControl(), controllers.UpdateProduct)
	product.Put("/delete-product/:id", middleware.TokenControl(), controllers.DeleteProduct)
	product.Put("/rebin-product/:id", middleware.TokenControl(), controllers.RebinProduct)
}
