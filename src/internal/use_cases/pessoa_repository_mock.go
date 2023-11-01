package usecases

type PessoaRepository struct {
}

func NewPessoaRepository() *PessoaRepository {
	return &PessoaRepository{}
}

func (r *PessoaRepository) GetContagem() (int64, error) {
	return 0, nil
}
