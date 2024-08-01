package main

import (
	"context"
	"log"
	"os"

	"github.com/TOMMy-Net/kafka-mess/db"
	"github.com/TOMMy-Net/kafka-mess/internal/handlers/api"
	"github.com/TOMMy-Net/kafka-mess/internal/kafka"
	"github.com/TOMMy-Net/kafka-mess/internal/routes"
	"github.com/TOMMy-Net/kafka-mess/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	kafka, err := kafka.ConnectBroker(os.Getenv("KAFKA_TOPIC"), 0)
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	go services.MessageKeeper(ctx, kafka, db) // keeper for un sended messages
	go kafka.ConnectKeeperV2()                // keeper for connect to kafka

	// init config for api
	server := api.NewApi()
	server.DB = db
	server.Kafka = kafka

	app := fiber.New(fiber.Config{
	
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, this is api service for messaging to kafka !")
	})

	routes.SetupRoutes(app, server)

	app.Listen(":8000")
}
