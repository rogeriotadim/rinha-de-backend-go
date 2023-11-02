package usecases

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rogeriotadim/rinha-de-backend-go/internal/infra/database"
	"github.com/stretchr/testify/assert"
)

func TestNewGetContagemUseCaseIntegration(t *testing.T) {
	// db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := database.NewPessoaRepository(db)
	uc := NewGetContagemUseCase(repo)
	assert.NotNil(t, uc)
	total, _ := uc.PessoaRepository.GetContagem()
	assert.Equal(t, int64(3), total)
}
