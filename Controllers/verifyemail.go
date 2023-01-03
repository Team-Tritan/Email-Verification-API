package controllers

import (
	"io/ioutil"
	"log"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Authkeys []string `yaml:"auth_keys"`
}

func VerifyEmail(ctx *fiber.Ctx) error {
	email := ctx.Params("email")
	user_auth := ctx.Query("token")
	config_file, config_err := ioutil.ReadFile("config.yaml")

	if config_err != nil {
		log.Fatalf("Config Error: %v", config_err)
	}

	config := &AppConfig{}
	parsing_err := yaml.Unmarshal(config_file, &config)
	if parsing_err != nil {
		log.Fatalf("Config Parsing Error: %v", parsing_err)
	}

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
		// do stuff
		return ctx.JSON(fiber.Map{
			"uwu?":  "uwu indeed",
			"email": email,
		})
	}
}
