package repository

import (
	"log"
	"testing"

	"github.com/google/uuid"

	"github.com/golang-crud/model"
	"github.com/golang-crud/repo"
)

var repository = PessoaRepository{}

func init() {

	err := repo.AbrirConexao()

	if err != nil {
		log.Fatal("Não foi possivel abrir a conexao com o banco")
	}

	repository.Db = repo.Db
}

func TestObterTodos(t *testing.T) {

	_, err := repository.ObterTodos()

	if err != nil {
		t.Errorf("Não foi possivel obter os registros")
	}
}

func TestInserir(t *testing.T) {

	id, err := uuid.NewUUID()

	if err != nil {
		t.Errorf("Não foi possivel gerar o uuid")
	}

	pessoa := model.Pessoa{}
	pessoa.ID = id.String()
	pessoa.Idade = 16
	pessoa.Nome = "go teste"
	pessoa.Telefone = "9999999999"

	err = repository.Incluir(pessoa)

	if err != nil {
		t.Errorf("Não foi possivel inserir o registro : %s", err.Error())
	}

}
func TestAlterar(t *testing.T) {

	pessoas, err := repository.ObterTodos()
	if err != nil {
		t.Errorf("Não foi possivel obter todas as pessoas")
	}
	if len(pessoas) == 0 {
		t.Errorf("Não existe pessoas cadastradas no banco")
	}

	pessoas[0].Nome = "teste go alterado"

	err = repository.Alterar(pessoas[0])
	if err != nil {
		t.Errorf("Não foi possivel alterar o registro : %s", err.Error())
	}
}

func TestDeletar(t *testing.T) {
	pessoas, err := repository.ObterTodos()

	if err != nil {
		t.Errorf("Não foi possivel obter todas as pessoas")
	}
	if len(pessoas) == 0 {
		t.Errorf("Não existe pessoas cadastradas no banco")
	}

	err = repository.Deletar(pessoas[0].ID)
	if err != nil {
		t.Errorf("Não foi possivel excluir o registro : %s", err.Error())
	}
}
