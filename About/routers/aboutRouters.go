package routers

import (
	"about/controllers"
	"about/middleware"

	"github.com/gofiber/fiber/v2"
)

func AboutRouters(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	about := v1.Group("/about")

	about.Post("/create-sector", middleware.TokenControl(), controllers.CreateAboutSector)
	about.Put("/update-sector/:id", middleware.TokenControl(), controllers.UpdateAboutSector)
	about.Delete("/delete-sector", middleware.TokenControl(), controllers.DeleteAboutSector)
	about.Get("/get-all-sectors", middleware.TokenControl(), controllers.GetAllSector)
	about.Get("/get-sector", middleware.TokenControl(), controllers.GetIDBySector)
}
