package routes

import (
	handlers "github.com/TOMMy-Net/kafka-mess/internal/handlers/api"
	"github.com/gofiber/fiber/v2"
)



func SetupMessRoutes(r fiber.Router, handler *handlers.ApiHandlers) {
	api := r.Group("/message")
	api.Get("/", handler.GetMessages)
	api.Post("/", handler.SendMessage)
	api.Get("/stats", handler.GetStatsMessages)
}
