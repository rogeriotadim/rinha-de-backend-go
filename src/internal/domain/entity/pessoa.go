package entity

import (
	"errors"
	"strconv"
	"strings"

	"github.com/rogeriotadim/rinha-de-backend-go/pkg"
)

type Pessoa struct {
	ID         string
	Apelido    string
	Nome       string
	Nascimento string
	Stack      []string
}

func NewPessoa(apelido, nome, nascimento string, stack []string) (pessoa *Pessoa, err error) {
	if apelido == "" {
		return nil, errors.New("apelido inválido")
	}

	if nome == "" {
		return nil, errors.New("nome inválido")
	}

	if nascimento == "" {
		return nil, errors.New("nascimento inválido")
	}

	data := strings.Split(nascimento, "-")
	ano, err := strconv.Atoi(data[0])
	if err != nil || ano < 1 || ano > 2023 {
		return nil, errors.New("ano inválido")
	}

	mes, err := strconv.Atoi(data[1])
	if err != nil || mes < 1 || mes > 12 {
		return nil, errors.New("mês inválido")
	}

	dia, err := strconv.Atoi(data[2])
	if err != nil || dia < 1 || dia > 31 {
		return nil, errors.New("dia inválido")
	}

	return &Pessoa{
		ID:         pkg.NewID().String(),
		Apelido:    apelido,
		Nome:       nome,
		Nascimento: nascimento,
		Stack:      stack,
	}, nil
}
