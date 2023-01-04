package controllers

import (
	"github.com/gofiber/fiber/v2"

	"tritan.dev/config"
	"tritan.dev/database"
)

func CheckStatus(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user_auth := ctx.Query("token")
	config := ctx.Locals("config").(*config.AppConfig)

	auth_keys := config.Authkeys
	isAuth := false
	for _, item := range auth_keys {
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

		db := database.New("./database/users.json")
		user := db.Get(id)

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
