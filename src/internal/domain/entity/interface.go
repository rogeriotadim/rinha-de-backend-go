package entity

type PessoaRepositoryInterface interface {
	GetContagem() (int64, error)
	// GetPessoaByID(id string) (pessoa Pessoa, err error)
	// GetPessoaByTermo(termo string) (pessoa Pessoa, err error)
	// CreatePessoa(p *Pessoa) (Pessoa, error)
}
