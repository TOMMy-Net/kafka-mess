package routes

import (
	"github.com/TOMMy-Net/kafka-mess/internal/handlers/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App, handler *api.ApiHandlers) {
	api := app.Group("/api", logger.New())

	SetupMessRoutes(api, handler)
}
