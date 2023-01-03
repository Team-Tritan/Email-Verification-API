package controllers

import (
	"github.com/gofiber/fiber/v2"
	"tritan.dev/config"
	"tritan.dev/mail"
)

func VerifyEmail(ctx *fiber.Ctx) error {
	email := ctx.Params("email")
	user_auth := ctx.Query("token")
	config := ctx.Locals("config").(*config.AppConfig)

	auth_keys := config.Authkeys
	isAuth := false
	for _, item := range auth_keys { // how do I loop through specific array in yaml instead of entire thing to match api key? essentially an .indludes()
		if item == user_auth {
			isAuth = true
			break
		}
	}

	if isAuth == false {
		return ctx.JSON(fiber.Map{
			"error":   true,
			"code":    403,
			"message": "You are not authenticated.",
		})
	} else {
		go mail.SendMail(email, ctx)

		return ctx.JSON(fiber.Map{
			"error":   false,
			"status":  200,
			"message": "Verification email sent successfully.",
			"address": email,
		})
	}
}
