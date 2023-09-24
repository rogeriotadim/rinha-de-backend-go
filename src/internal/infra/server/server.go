package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	ph := NewPessoaHandler()
	ch := NewContagemHandler()

	config := fiber.Config{
		ServerHeader: "rinha de backend", // add custom server header
	}
	app := fiber.New(config)
	app.Post("/pessoas/", ph.CreatePessoa)
	app.Get("/pessoas/:id", ph.GetPessoaByID)
	app.Get("/pessoas/", ph.GetPessoaByTerm)

	// app.Get("/contagem-pessoas", func(c *fiber.Ctx) error {
	// 	return c.SendString("0")
	// })
	app.Get("/contagem-pessoas", ch.ContagemPessoas)

	log.Fatal(app.Listen(":3000"))
}
