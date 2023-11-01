package database

import (
	"database/sql"

	"github.com/rogeriotadim/rinha-de-backend-go/internal/domain/entity"
)

type PessoaRepository struct {
	Db *sql.DB
}

func NewPessoaRepository(db *sql.DB) *PessoaRepository {
	return &PessoaRepository{Db: db}
}

func (r *PessoaRepository) CreatePessoa(pessoa *entity.Pessoa) error {
	stmt, err := r.Db.Prepare("INSERT INTO PESSOAS (ID, APELIDO, NOME, NASCIMENTO, STACK) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(pessoa.ID, pessoa.Apelido, pessoa.Nome, pessoa.Nascimento, pessoa.Stack)
	if err != nil {
		return err
	}
	return nil
}

func (r *PessoaRepository) GetContagem() (int64, error) {
	var total int64
	err := r.Db.QueryRow("Select count(*) from pessoas").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
