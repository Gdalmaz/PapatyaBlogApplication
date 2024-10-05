package routers

import (
	"panel/controllers"
	"panel/middleware"

	"github.com/gofiber/fiber/v2"
)

func PanelRouters(app *fiber.App)  {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	panel := v1.Group("/panel")

	panel.Get("/get-deleted-product", middleware.TokenControl(), controllers.GetDeletedProducts)
	panel.Get("/get-avaible-product", middleware.TokenControl(), controllers.GetAvaibleProducts)
}
