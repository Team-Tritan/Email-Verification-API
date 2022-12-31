package router

import (
	"github.com/gofiber/fiber"

	Controllers "tritan.dev/Controllers"
)


func BuildAPI(app *fiber.App) error {
	app.Get("/", Controllers.ServeRoot)

	api := app.Group("/api")
	api.Post("/verify", Controllers.VerifyEmail)
	api.Post("/check", Controllers.CheckStatus)

	api.Get("/status", func(ctx *fiber.Ctx) {
		ctx.JSON(fiber.Map{
			"error": false,
			"code": 200,
			"message": "more alive than u", 
		})
	})

	return nil
}


func HandleRoutes(app *fiber.App) error { 
	BuildAPI(app)

	return nil
}


