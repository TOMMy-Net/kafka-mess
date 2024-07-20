package main

import (
	"github.com/TOMMy-Net/kafka-mess/internal/routes"
	"github.com/gofiber/fiber/v2"
)
 
func main() {
   app := fiber.New()
   routes.SetupRoutes(app)

   
   app.Listen(":8000")
}
