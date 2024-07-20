package api

import (
	handlers "github.com/TOMMy-Net/kafka-mess/internal/handlers/api"
	"github.com/gofiber/fiber/v2")

	
func SetupMessRoutes(r fiber.Router) {
	api := r.Group("/message")
	api.Get("/", handlers.GetMessage)
	api.Post("/", handlers.SetMessage)
}