package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"tritan.dev/database"
)

type User struct {
	Email     string `json:"email"`
	Verified  string `json:"verified"`
	Datesent  string `json:"date_sent"`
	Verifcode string `json:"verif_code"`
}

func VerifyEmail(ctx *fiber.Ctx) error {
	user_token := ctx.Params("id")
	db := database.New("./database/users.json")

	user := db.Get(user_token)
	fmt.Print(user)

	if user == nil {
		return ctx.JSON(fiber.Map{
			"error":   true,
			"code":    404,
			"message": "Verification token not found.",
		})
	}

	// How do I map SPECIFIC data from the db to the struct? or get from the existing db and resave?
	// user := User{}
	// user.Email = user_exists["email"]
	// user.Datesent = user_exists["date_sent"]
	// user.Verifcode = user_exists["verif_code"]

	db.Set(user_token, map[string]interface{}{
		//"email": user_exists["Email"],
		//"date_sent":  user.Datesent,
		"verified":   true,
		"verif_code": user_token,
	})

	db.Save("./database/users.json")

	finalRecord := db.Get(user_token)

	return ctx.JSON(finalRecord)
}
