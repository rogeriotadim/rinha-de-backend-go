package main

import (
	"context"

	"github.com/rogeriotadim/rinha-de-backend-go/cmd/configs"
	"github.com/rogeriotadim/rinha-de-backend-go/internal/infra/database"

	// mysql
	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()
	configs, err := configs.LoadConfig(".")
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
	rows, err := db.Query(ctx, "SELECT count(id) as qtd FROM pessoas")
	if err != nil {
		panic(err)
	}
	rows.Next()
	var qtd *int64
	if err := rows.Scan(&qtd); err != nil {
		panic(err)
	}

	println("Vai CORINTHIANS!!!")
	println(*qtd)
	// server.Run()
}
