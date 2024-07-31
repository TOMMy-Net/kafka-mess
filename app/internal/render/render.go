package render

import "github.com/gofiber/fiber/v2"

var (
	GoodMsg = "the message was successfully delivered"
)

type Answer struct {
	Status  string `json:"status"`
	Message string `json:"msg"`
}

func SendServerError(c *fiber.Ctx, err error) error{
	return c.Status(fiber.StatusInternalServerError).JSON(Error{
		Status: "error",	
		Error: err.Error(),	
	})
}