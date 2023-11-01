package usecases

import "github.com/rogeriotadim/rinha-de-backend-go/internal/domain/entity"

type GetContagemUseCase struct {
	PessoaRepository entity.PessoaRepositoryInterface
}

func NewGetContagemUseCase(pessoaRepository entity.PessoaRepositoryInterface) *GetContagemUseCase {
	return &GetContagemUseCase{
		PessoaRepository: pessoaRepository,
	}
}
