package router

import (
	"github.com/gofiber/fiber"

	controllers "tritan.dev/controllers"
)


func Build_Root(app *fiber.App) error {
	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.JSON(fiber.Map{
			"error": false,
			"code": 200,
			"hi": "fuck off :)",
		})
	})

	return nil
}

func Build_API(app *fiber.App) error {
	api := app.Group("/api")

	api.Post("/send", controllers.VerifyEmail)
	api.Post("/check", controllers.CheckStatus)

	app.Get("/api/status", func(ctx *fiber.Ctx) {
		ctx.JSON(fiber.Map{
			"error": false,
			"code": 200,
			"message": "more alive than u", 
		})
	})

	return nil
}


func ServeRoutes(app *fiber.App) error { 
	Build_Root(app)
	Build_API(app)

	return nil
}


