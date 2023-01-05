package router

import (
	"github.com/gofiber/fiber/v2"

	"tritan.dev/controllers"
)

func BuildAPI(app *fiber.App) error {
	app.Get("/", controllers.ServeRoot)

	api := app.Group("/api")
	api.All("/send/:email", controllers.VerifyEmail)
	api.All("/check/:id", controllers.CheckStatus)

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
