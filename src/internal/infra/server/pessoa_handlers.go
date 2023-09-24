package server

import (
	"github.com/gofiber/fiber/v2"
)

type PessoaHandler struct {
}

func NewPessoaHandler() *PessoaHandler {
	return &PessoaHandler{}
}

func (p *PessoaHandler) CreatePessoa(c *fiber.Ctx) error {
	return c.
		Status(fiber.StatusInternalServerError).
		JSON(fiber.Map{"error": "internal server error"})
}

func (p *PessoaHandler) GetPessoaByID(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "404"})

}

func (p *PessoaHandler) GetPessoaByTerm(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "404"})

}
