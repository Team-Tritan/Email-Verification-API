package controllers

import (
	"github.com/gofiber/fiber/v2"

	"tritan.dev/config"
	"tritan.dev/database"
)

func CheckStatus(ctx *fiber.Ctx) error {
	user_token := ctx.Params("id")
	api_key := ctx.Query("token")
	config := ctx.Locals("config").(*config.AppConfig)

	auth_keys := config.Authkeys
	isAuth := false
	for _, item := range auth_keys {
		if item == api_key {
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

		db := database.New("./database/users.json")
		user := db.Get(user_token)

		if user == nil {
			return ctx.JSON(fiber.Map{
				"error":   true,
				"code":    404,
				"message": "ID provided not found.",
			})
		}

		return ctx.JSON(user)
	}
}
