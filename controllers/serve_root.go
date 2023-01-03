package controllers

import "github.com/gofiber/fiber/v2"

func ServeRoot(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"error":   false,
		"code":    200,
		"message": "fuck off :)",
	})
}
