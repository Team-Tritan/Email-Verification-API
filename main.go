package main

import (
	"log"

	"github.com/gofiber/fiber"
	router "tritan.dev/Router"
)


func main() { 
  app := fiber.New()  

  err := router.HandleRoutes(app)   

  app.Listen(":3000") 

  if err != nil {
	log.Fatalf("Error: %v", err)
  } 
}