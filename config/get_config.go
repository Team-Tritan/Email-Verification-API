package config

import (
	"io/ioutil"
	"log"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Authkeys []string `yaml:"auth_keys"`

	Mail struct {
		FromAddress string `yaml:"from_address"`
		MailServer  string `yaml:"mail_server"`
		Username    string `yaml:"username"`
		Password    string `yaml:"password"`
		Port        int    `yaml:"port"`
		Tls         bool   `yaml:"tls"`
	} `yaml:"mail"`
}

func (config *AppConfig) LoadConfig(app *fiber.App) error {

	app.Use(func(ctx *fiber.Ctx) {
		config_file, config_err := ioutil.ReadFile("config.yaml")

		if config_err != nil {
			log.Fatalf("Config Error: %v", config_err)
		}

		parsing_err := yaml.Unmarshal(config_file, &config)
		if parsing_err != nil {
			log.Fatalf("Config Parsing Error: %v", parsing_err)
		}

		ctx.Locals("config", config)
	})

	return nil
}
