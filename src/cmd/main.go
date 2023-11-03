package main

import (
	"context"
	"fmt"
	"os"

	"github.com/rogeriotadim/rinha-de-backend-go/cmd/configs"
	"github.com/rogeriotadim/rinha-de-backend-go/internal/infra/database"
	usecases "github.com/rogeriotadim/rinha-de-backend-go/internal/use_cases"

	// mysql
	_ "github.com/lib/pq"
)

func main() {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	fmt.Println("########### >>>>>>>>>>>> ", currentDir)

	ctx := context.Background()
	configs, err := configs.LoadConfig("/app")
	if err != nil {
		panic(err)
	}
	// connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName)
	// db, err := sql.Open(configs.DBDriver, connStr)
	db, err := database.NewDatabasePool(configs, ctx)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repoPessoas := database.NewPessoaRepository(ctx, db)
	ucContagem := usecases.NewGetContagemUseCase(repoPessoas)

	qtd, err := ucContagem.PessoaRepository.GetContagem()
	if err != nil {
		panic(err)
	}

	println("Vai CORINTHIANS!!!")
	println(qtd)
	// server.Run()
}
