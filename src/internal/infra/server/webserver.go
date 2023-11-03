package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type WebServer struct {
	App *fiber.App
}

func NewWebServer(app *fiber.App) *WebServer {
	ch := NewContagemHandler()

	app.Get("/contagem-pessoas", ch.ContagemPessoas)
	return &WebServer{App: app}
}

func Run() {
	ph := NewPessoaHandler()

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

	log.Fatal(app.Listen(":3000"))
}
