package main

import (
	"fmt"
	"log"
	"os"

	"github.com/TOMMy-Net/kafka-mess/db"
	"github.com/TOMMy-Net/kafka-mess/internal/handlers/api"
	"github.com/TOMMy-Net/kafka-mess/internal/kafka"
	"github.com/TOMMy-Net/kafka-mess/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	//ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	kafka, err := kafka.Connect([]string{fmt.Sprintf("%s:%s",os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT"))}, os.Getenv("KAFKA_TOPIC"))
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	//go services.MessageKeeper(ctx, kafka, db) // keeper for un sended messages

	// init config for api
	server := api.NewApi()
	server.DB = db
	server.Kafka = kafka

	app := fiber.New(fiber.Config{})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, this is api service for messaging to kafka !")
	})

	routes.SetupRoutes(app, server)

	app.Listen(":8000")
}
