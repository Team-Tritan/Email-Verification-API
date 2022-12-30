package main

import (
	"log"
	router "tritan.dev/v/router"
	"github.com/gofiber/fiber"
)


func main() { 
  app := fiber.New()  

  router_err := router.mainRouter(app) // WHY
  
  if router_err != nil{
	log.Fatalf("Error: %v", router_err)
  } 

  listen_err := app.Listen(":3000") 

  if listen_err != nil{
	log.Fatalf("Error: %v", router_err)
  } 
}