package database

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rogeriotadim/rinha-de-backend-go/cmd/configs"
)

var (
	db   *pgxpool.Pool
	once sync.Once
)

func NewDatabasePool(config *configs.Conf, ctx context.Context) (*pgxpool.Pool, error) {
	var err error = nil

	// config, err := configs.LoadConfig(".")
	// if err != nil {
	// 	return nil, err
	// }

	once.Do(func() {
		connUrl := fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			config.DBUser,
			config.DBPassword,
			config.DBHost,
			config.DBPort,
			config.DBName,
		)

		poolConfig, err := pgxpool.ParseConfig(connUrl)
		if err != nil {
			log.Fatalln("Unable to parse connection url:", err)
		}

		db, err = pgxpool.NewWithConfig(ctx, poolConfig)
		if err != nil {
			log.Fatalln("Unable to create connection pool:", err)
		}

		if err := db.Ping(ctx); err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
	})

	return db, err

}
