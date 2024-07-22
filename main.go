package main

import (
	"log"

	"github.com/TOMMy-Net/kafka-mess/internal/routes"
   "github.com/TOMMy-Net/kafka-mess/db"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)
 
func main() {
   err := godotenv.Load()
   if err != nil{
      log.Fatal(err)
   }

   if err := db.Connect(); err != nil {
      log.Fatal(err)
   }

   app := fiber.New()
   routes.SetupRoutes(app)

   
   app.Listen(":8000")
}
