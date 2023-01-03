package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gofiber/fiber"
	"gopkg.in/yaml.v3"

	router "tritan.dev/Router"
)

func main() {
	fmt.Println("Starting!")
	app := fiber.New()
	config_file, config_err := ioutil.ReadFile("config.yaml")

	if config_err != nil {
		log.Fatalf("Config Error: %v", config_err)
	}

	config := make(map[interface{}]interface{})
	parsing_err := yaml.Unmarshal(config_file, &config)
	if parsing_err != nil {
		log.Fatalf("Config Parsing Error: %v", parsing_err)
	}

	fmt.Println("Loading config...")
	for k, v := range config {
		fmt.Printf("%s -> %d\n", k, v)
	}

	router_err := router.HandleRoutes(app)
	if router_err != nil {
		log.Fatalf("Router Error: %v", router_err)
	}

	listen_err := app.Listen(config["port"])
	if listen_err != nil {
		log.Fatalf("Listening Error: %v", listen_err)
	}
}
