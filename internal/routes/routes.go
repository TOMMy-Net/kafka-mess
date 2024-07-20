package routes

import (
	apiRoutes "github.com/TOMMy-Net/kafka-mess/internal/routes/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	apiRoutes.SetupMessRoutes(api)

}
