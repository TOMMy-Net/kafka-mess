package api

import (
	"fmt"

	"github.com/TOMMy-Net/kafka-mess/db"
	"github.com/TOMMy-Net/kafka-mess/internal/kafka"
	"github.com/TOMMy-Net/kafka-mess/internal/models"
	"github.com/TOMMy-Net/kafka-mess/internal/render"
	"github.com/gofiber/fiber/v2"
)

type ApiHandlers struct {
	DB    *db.Database  // field for databse driver
	Kafka *kafka.Broker // field for kafka driver
}

func NewApi() *ApiHandlers {
	return &ApiHandlers{}
}

func (a *ApiHandlers) GetMessages(c *fiber.Ctx) error {
	mess, err := a.Kafka.Read(c.Context())
	if err != nil {
		render.SendServerError(c, kafka.ErrWithRead)
	}

	return c.Status(fiber.StatusOK).JSON(render.Answer{
		Status:  "ok",
		Message: fmt.Sprintf("Key: %s | Message: %s", mess.Key, mess.Value),
	})

}


func (a *ApiHandlers) SendMessage(c *fiber.Ctx) error {
	var mess models.Message

	if err := c.BodyParser(&mess); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(render.Error{
			Status: "error",
			Error:  err.Error(),
		})
	}

	if err := models.ValidStruct(&mess); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(render.Error{
			Status: "error",
			Error:  err.Error(),
		})
	}

	id, err := a.DB.WriteMessage(c.Context(), mess)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(render.Error{
			Status: "error",
			Error:  db.ErrBaseWrite.Error(),
		})
	}
	mess.UID = id.String()

	err = a.Kafka.SendMessage(mess)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(render.Answer{
			Status:  "ok",
			Message: kafka.ErrWithWrite.Error(),
		})
	}

	a.DB.UpdateMessageStatus(c.Context(), mess.UID, 1)

	return c.Status(fiber.StatusOK).JSON(render.Answer{
		Status:  "ok",
		Message: render.GoodMsg,
	})
}
