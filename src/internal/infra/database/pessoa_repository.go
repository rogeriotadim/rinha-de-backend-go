package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rogeriotadim/rinha-de-backend-go/internal/domain/entity"
)

type PessoaRepository struct {
	Db  *pgxpool.Pool
	Ctx context.Context
}

func NewPessoaRepository(ctx context.Context, db *pgxpool.Pool) *PessoaRepository {
	return &PessoaRepository{
		Ctx: ctx,
		Db:  db,
	}
}

func (r *PessoaRepository) CreatePessoa(pessoa *entity.Pessoa) error {
	_, err := r.Db.Query(r.Ctx, "INSERT INTO PESSOAS (ID, APELIDO, NOME, NASCIMENTO, STACK) VALUES (?, ?, ?, ?, ?)", pessoa.ID, pessoa.Apelido, pessoa.Nome, pessoa.Nascimento, pessoa.Stack)
	if err != nil {
		return err
	}

	return nil
}

func (r *PessoaRepository) GetContagem() (int64, error) {
	var total int64
	err := r.Db.QueryRow(r.Ctx, "Select count(*) from pessoas").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
