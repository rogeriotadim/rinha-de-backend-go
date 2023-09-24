package entity

import (
	"testing"

	"github.com/rogeriotadim/rinha-de-backend-go/pkg"
	"github.com/stretchr/testify/assert"
)

func TestNewPessoa(t *testing.T) {
	stack := []string{"Go", "Kubernetes", "AWS"}
	pessoa, err := NewPessoa("Doe", "John", "1971-12-25", stack)
	assert.Nil(t, err)
	assert.NotNil(t, pessoa)
	assert.NotEmpty(t, pessoa.ID)
	assert.NotEmpty(t, pessoa.Apelido)
	assert.NotEmpty(t, pessoa.Nome)
	assert.NotEmpty(t, pessoa.Nascimento)
	assert.Equal(t, "Doe", pessoa.Apelido)
	assert.Equal(t, "John", pessoa.Nome)
	assert.Equal(t, "1971-12-25", pessoa.Nascimento)
	assert.Equal(t, stack, pessoa.Stack)
	_, err = pkg.ParseID(pessoa.ID)
	assert.Nil(t, err)
}

func TestNewPessoaNoApelido(t *testing.T) {
	stack := []string{"Go", "Kubernetes", "AWS"}
	_, err := NewPessoa("", "John", "1971-12-25", stack)
	assert.EqualError(t, err, "apelido inválido")
}

func TestNewPessoaNoName(t *testing.T) {
	stack := []string{"Go", "Kubernetes", "AWS"}
	_, err := NewPessoa("Doe", "", "1971-12-25", stack)
	assert.EqualError(t, err, "nome inválido")
}

func TestNewPessoaNoNascimento(t *testing.T) {
	stack := []string{"Go", "Kubernetes", "AWS"}
	_, err := NewPessoa("Doe", "John", "", stack)
	assert.EqualError(t, err, "nascimento inválido")
}

func TestNewPessoaErrorAnoMin(t *testing.T) {
	stack := []string{"Go", "Kubernetes", "AWS"}
	_, err := NewPessoa("Doe", "John", "0000-12-25", stack)
	assert.EqualError(t, err, "ano inválido")
}

func TestNewPessoaErrorAnoMax(t *testing.T) {
	stack := []string{"Go", "Kubernetes", "AWS"}
	_, err := NewPessoa("Doe", "John", "2024-12-25", stack)
	assert.EqualError(t, err, "ano inválido")
}

func TestNewPessoaErrorMesMin(t *testing.T) {
	stack := []string{"Go", "Kubernetes", "AWS"}
	_, err := NewPessoa("Doe", "John", "1971-0-25", stack)
	assert.EqualError(t, err, "mês inválido")
}
func TestNewPessoaErrorMesMax(t *testing.T) {
	stack := []string{"Go", "Kubernetes", "AWS"}
	_, err := NewPessoa("Doe", "John", "1971-13-25", stack)
	assert.EqualError(t, err, "mês inválido")
}

func TestNewPessoaErrorDiaMin(t *testing.T) {
	stack := []string{"Go", "Kubernetes", "AWS"}
	_, err := NewPessoa("Doe", "John", "1971-12-0", stack)
	assert.EqualError(t, err, "dia inválido")
}

func TestNewPessoaErrorDiaMax(t *testing.T) {
	stack := []string{"Go", "Kubernetes", "AWS"}
	_, err := NewPessoa("Doe", "John", "1971-12-32", stack)
	assert.EqualError(t, err, "dia inválido")
}
