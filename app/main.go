package main

import (
	"log"

	"github.com/TOMMy-Net/kafka-mess/db"
	"github.com/TOMMy-Net/kafka-mess/internal/kafka"
	"github.com/TOMMy-Net/kafka-mess/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	if err := kafka.ConnectBroker(); err != nil {
		log.Fatal(err)
	}

	
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, this is api service for messaging to kafka !")
	})

	routes.SetupRoutes(app)

	app.Listen(":8000")
}
