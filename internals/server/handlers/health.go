package handlers

import "github.com/gofiber/fiber/v3"

func HelloWorldHandler(c fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}

	return c.JSON(resp)
}
