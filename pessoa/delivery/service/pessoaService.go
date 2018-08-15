package service

import (
	"github.com/golang-crud/model"
	"github.com/golang-crud/pessoa/delivery/repository"
	"github.com/golang-crud/repo"
)

//PessoaService realiza servicos de pessoa
type PessoaService struct {
	Repository repository.PessoaRepository
}

//InitRepository inicia os repository
func InitRepository(service *PessoaService) (pessoaService *PessoaService) {
	service.Repository = repository.PessoaRepository{}
	service.Repository.Db = repo.Db

	return service
}

//ObterPessoas obtem todas as pessoas do banco
func (service *PessoaService) ObterPessoas() (pessoas []model.Pessoa, err error) {

	service = InitRepository(service)

	pessoas, err = service.Repository.ObterTodos()

	if err != nil {
		return nil, err
	}

	return pessoas, err
}
