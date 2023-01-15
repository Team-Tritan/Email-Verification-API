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
	id := ctx.Params("id")
	db := database.New("./database/users.json")

	user := db.Get(id)
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

	db.Set(id, map[string]interface{}{
		//"email": user_exists["Email"],
		//"date_sent":  user.Datesent,
		"verified":   true,
		"verif_code": id,
	})

	db.Save("./database/users.json")

	finalRecord := db.Get(id)

	return ctx.JSON(finalRecord)
}
