package main

import (
	"log"

	"github.com/gofiber/fiber"
	router "tritan.dev/router"
)


func main() { 
  app := fiber.New()  

  err := router.ServeRoutes(app)   

  app.Listen(":3000") 

  if err != nil {
	log.Fatalf("Error: %v", err)
  } 
}