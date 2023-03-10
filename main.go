package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"tritan.dev/config"
	"tritan.dev/router"
)

func main() {
	fmt.Println("Starting!")
	app := fiber.New()

	appConfig := config.AppConfig{}
	appConfig.LoadConfig(app)

	env_err := godotenv.Load()
	if env_err != nil {
		log.Fatalf("ENV Error: %v", env_err)
	}

	port := os.Getenv("PORT")

	router_err := router.HandleRoutes(app)
	if router_err != nil {
		log.Fatalf("Router Error: %v", router_err)
	}

	listen_err := app.Listen(port)
	if listen_err != nil {
		log.Fatalf("Listening Error: %v", listen_err)
	}
}
