package usecases

import (
	"context"
	"testing"

	"github.com/rogeriotadim/rinha-de-backend-go/cmd/configs"
	"github.com/rogeriotadim/rinha-de-backend-go/internal/infra/database"
	"github.com/stretchr/testify/assert"
)

func TestNewGetContagemUseCaseIntegration(t *testing.T) {
	ctx := context.Background()
	configs, err := configs.LoadConfig("/app")
	if err != nil {
		panic(err)
	}
	// connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName)
	// db, err := sql.Open(configs.DBDriver, connStr)
	pool, err := database.NewDatabasePool(configs, ctx)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	repo := database.NewPessoaRepository(context.Background(), pool)
	uc := NewGetContagemUseCase(repo)
	assert.NotNil(t, uc)
	total, _ := uc.PessoaRepository.GetContagem()
	assert.GreaterOrEqual(t, int64(0), total)
}
