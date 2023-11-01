package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGetContagemUseCase(t *testing.T) {
	repo := NewPessoaRepository()
	uc := NewGetContagemUseCase(repo)
	assert.NotNil(t, uc)
	total, _ := uc.PessoaRepository.GetContagem()
	assert.Equal(t, int64(1), total)
}
