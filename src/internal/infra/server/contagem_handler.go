package server

import "github.com/gofiber/fiber/v2"

type ContagemHandler struct {
}

func NewContagemHandler() *ContagemHandler {
	return &ContagemHandler{}
}

func (cp *ContagemHandler) ContagemPessoas(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "404"})

}
