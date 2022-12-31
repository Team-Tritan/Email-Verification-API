package controllers

import "github.com/gofiber/fiber"

func ServeRoot(ctx *fiber.Ctx) {	
	ctx.JSON(fiber.Map{
		"error": false,
		"code": 200,
		"message": "fuck off :)",
	})
}