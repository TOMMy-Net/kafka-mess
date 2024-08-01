package api

import (
	"github.com/TOMMy-Net/kafka-mess/internal/models"
	"github.com/TOMMy-Net/kafka-mess/internal/render"
	"github.com/gofiber/fiber/v2"
)

func (a *ApiHandlers) GetStatsMessages(c *fiber.Ctx) error {
	var stats models.MessageStats
	count, err := a.DB.CountMessages(c.Context())
	if err != nil {
		return render.SendServerError(c, err)
	}

	unSend, err := a.DB.CountUnSend(c.Context())
	if err != nil {
		return render.SendServerError(c, err)
	}

	stats.TotalMessages = count
	stats.UnSendMessages = unSend
	stats.SendMessages = stats.TotalMessages - stats.UnSendMessages

	return c.Status(fiber.StatusOK).JSON(stats)
}
