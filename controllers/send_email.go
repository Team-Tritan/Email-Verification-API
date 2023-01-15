package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"tritan.dev/config"
	"tritan.dev/database"
	"tritan.dev/mail"
)

func SendEmail(ctx *fiber.Ctx) error {
	email := ctx.Params("email")
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
		date := time.Now()
		userkey := uuid.New().String()

		go mail.SendMail(email, userkey, ctx)

		db := database.New("./database/users.json")
		db.Set(userkey, map[string]interface{}{
			"email":      email,
			"verified":   false,
			"date_sent":  date.String(),
			"verif_code": userkey,
		})
		db.Save("./database/users.json")

		return ctx.JSON(fiber.Map{
			"status":  200,
			"message": "Verification email sent successfully.",
			"email":   email,
			"id":      userkey,
		})
	}
}
