package api

import (
	"github.com/TOMMy-Net/kafka-mess/db"
	"github.com/TOMMy-Net/kafka-mess/internal/models"
	"github.com/TOMMy-Net/kafka-mess/internal/render"
	"github.com/gofiber/fiber/v2"
)

func GetMessages(c *fiber.Ctx) error {

	return nil
}

func SendMessage(c *fiber.Ctx) error {
	m := new(models.Message)

	if err := c.BodyParser(m); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(render.Error{
			Status: "error",
			Error: err.Error(),
		})
	}

	if err := models.ValidStruct(m); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(render.Error{
			Status: "error",
			Error: err.Error(),
		})
	}
	
	_, err := db.WriteMessage(c.Context(), *m)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(render.Error{
			Status: "error",
			Error:  db.ErrBaseWrite.Error(),
		})
		
	}

	return c.Status(fiber.StatusOK).JSON(render.Answer{
		Status:  "ok",
		Message: render.GoodMsg,
	})
}
