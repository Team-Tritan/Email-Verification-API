package router

import (
	"github.com/gofiber/fiber/v2"

	Controllers "tritan.dev/Controllers"
)

func BuildAPI(app *fiber.App) error {
	app.Get("/", Controllers.ServeRoot)

	api := app.Group("/api")
	api.All("/verify/:email", Controllers.VerifyEmail)
	api.All("/check/:email", Controllers.CheckStatus)

	api.Get("/status", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"error":   false,
			"code":    200,
			"message": "more alive than u",
		})
	})

	return nil
}

func HandleRoutes(app *fiber.App) error {
	BuildAPI(app)

	return nil
}
