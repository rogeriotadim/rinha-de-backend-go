package entity

import (
	"errors"
	"fmt"
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
	err = validate(apelido, nome, nascimento, stack)
	if err != nil {
		return nil, err
	}

	return &Pessoa{
		ID:         pkg.NewID().String(),
		Apelido:    apelido,
		Nome:       nome,
		Nascimento: nascimento,
		Stack:      stack,
	}, nil
}

func validate(apelido, nome, nascimento string, stack []string) error {
	tipo, err := fmt.Printf("%T", apelido)
	if err != nil || apelido == "" || tipo != 6 {
		return errors.New("apelido inválido")
	}

	tipo, err = fmt.Printf("%T", nome)
	if err != nil || nome == "" || tipo != 6 {
		return errors.New("nome inválido")
	}

	tipo, err = fmt.Printf("%T", nascimento)
	if err != nil || nascimento == "" || tipo != 6 {
		return errors.New("nascimento inválido")
	}

	data := strings.Split(nascimento, "-")
	ano, err := strconv.Atoi(data[0])
	if err != nil || ano < 1 || ano > 2023 {
		return errors.New("ano inválido")
	}

	mes, err := strconv.Atoi(data[1])
	if err != nil || mes < 1 || mes > 12 {
		return errors.New("mês inválido")
	}

	dia, err := strconv.Atoi(data[2])
	if err != nil || dia < 1 || dia > 31 {
		return errors.New("dia inválido")
	}

	if len(stack) > 0 {
		for _, v := range stack {
			tipo, err = fmt.Printf("%T", v)
			if err != nil || tipo != 6 {
				return errors.New("stack inválida")
			}
		}
	}

	return nil
}
