package controllers

import (
	"io/ioutil"
	"log"

	"github.com/gofiber/fiber"
	"gopkg.in/yaml.v3"
)

type Query struct {
	email string `query:"email"`
}
type Config struct {
	auth_keys []string `yaml:"auth_keys"`
}

func CheckStatus(ctx *fiber.Ctx) {
	email := ctx.Params("email")
	user_auth := ctx.Query("token")
	config_file, config_err := ioutil.ReadFile("config.yaml")
	if config_err != nil {
		log.Fatalf("Config Error: %v", config_err)
	}

	var config Config
	parsing_err := yaml.Unmarshal(config_file, &config)
	if parsing_err != nil {
		log.Fatalf("Config Parsing Error: %v", parsing_err)
	}

	isAuth := false
	for _, item := range config.auth_keys {
		if item == user_auth {
			isAuth = true
			break
		}
	}

	if isAuth == false {
		ctx.JSON(fiber.Map{
			"error":   true,
			"code":    403,
			"message": "You are not authenticated.",
		})
	} else {
		// do stuff
		ctx.JSON(fiber.Map{
			"code":  200,
			"email": email,
		})
	}
}
