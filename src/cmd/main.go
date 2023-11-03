package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/rogeriotadim/rinha-de-backend-go/cmd/configs"
	"github.com/rogeriotadim/rinha-de-backend-go/internal/infra/database"
	"github.com/rogeriotadim/rinha-de-backend-go/internal/infra/server"
	usecases "github.com/rogeriotadim/rinha-de-backend-go/internal/use_cases"

	// mysql
	_ "github.com/lib/pq"
)

func main() {

	ctx := context.Background()
	configs, err := configs.LoadConfig("/app")
	if err != nil {
		panic(err)
	}
	db, err := database.NewDatabasePool(configs, ctx)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repoPessoas := database.NewPessoaRepository(ctx, db)
	ucContagem := usecases.NewGetContagemUseCase(repoPessoas)

	webServerConfig := fiber.Config{
		ServerHeader: "rinha de backend", // add custom server header
	}
	app := fiber.New(webServerConfig)
	webServer := server.NewWebServer(app)
	webServer.App.Listen(configs.WebServerPort)

	qtd, err := ucContagem.PessoaRepository.GetContagem()
	if err != nil {
		panic(err)
	}

	println("Vai CORINTHIANS!!!")
	println(qtd)
	// server.Run()
}
