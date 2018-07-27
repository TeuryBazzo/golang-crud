package httphandler

import (
	"encoding/json"
	"fmt"
	"go-clean-arch/article/delivery/http"
	"io/ioutil"
	"log"
	"projects/golang-crud/model"
	"projects/golang-crud/repo"
	"strconv"
)

//NewPessoaHTTPHandler redireciona o request do endpoint /pessoa
func NewPessoaHTTPHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		ObterPessoas(w, r)
	case "POST":
		CriarPessoa(w, r)
	case "PUT":
		AlterarPessoa(w, r)
	case "DELETE":
		DeletarPessoa(w, r)
	}

}

// ObterPessoas funcao que obtem pessoas
func ObterPessoas(w http.ResponseWriter, r *http.Request) {

	pessoa := model.Pessoa{}

	rows, err := repo.Db.Queryx("select * from Pessoa")

	if err != nil {
		http.Error(w, "Não encontramos nemhum item", http.StatusInternalServerError)
		fmt.Println("[main] [select pessoa] ", err.Error())
		return
	}

	for rows.Next() {
		err = rows.StructScan(&pessoa)
		if err != nil {
			http.Error(w, "Não encontramos nemhum item", http.StatusInternalServerError)
			fmt.Println("[main] [select pessoa] ", err.Error())
			return
		}
	}

	fmt.Fprint(w, pessoa)
	return
}

//CriarPessoa funcao que cria um registro na tabela pessoa
func CriarPessoa(w http.ResponseWriter, r *http.Request) {
	io, err := ioutil.ReadAll(r.Body)

	defer r.Body.Close()

	if err != nil {
		fmt.Fprint(w, "Error ao ler o body "+err.Error(), http.StatusBadRequest)
		return
	}

	var pessoa model.Pessoa

	err = json.Unmarshal(io, &pessoa)

	if err != nil {
		http.Error(w, "Error ao ler json "+err.Error(), http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf("INSERT INTO pessoa (id,name,age,telephone) VALUES (%d,'%s',%d,'%s')", pessoa.Id, pessoa.Nome, pessoa.Idade, pessoa.Telefone)

	fmt.Println(query)

	_, err = repo.Db.Queryx(query)

	if err != nil {
		http.Error(w, "Error ao inserir registro "+err.Error(), http.StatusBadRequest)
		log.Panicf("Error ao inserir registro Pessoa %s", err.Error())
		return
	}

	fmt.Fprint(w, "registro inserido")
	return
}

//AlterarPessoa funcao que altera um registro de pessoa
func AlterarPessoa(w http.ResponseWriter, r *http.Request) {

	io, err := ioutil.ReadAll(r.Body)

	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Conteudo vazio", http.StatusBadRequest)
		return
	}

	var pessoa model.Pessoa
	err = json.Unmarshal(io, &pessoa)

	if err != nil {
		http.Error(w, "Erro ao serializar json", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf("UPDATE pessoa SET name = '%s', age = %d, telephone = '%s' WHERE id = %d", pessoa.Nome, pessoa.Idade, pessoa.Telefone, pessoa.Id)

	_, err = repo.Db.Queryx(query)

	if err != nil {
		http.Error(w, "Erro ao alterar pessoa", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Registro alterado")
	return
}

//DeletarPessoa funcão que deleta um registro de pessoa do banco
func DeletarPessoa(w http.ResponseWriter, r *http.Request) {

	queryString := r.URL.Query()

	pessoaID := queryString["id"][0]

	ID, err := strconv.ParseInt(pessoaID, 10, 64)

	if err != nil {
		http.Error(w, "Pessoa id não é valido", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf("DELETE FROM pessoa WHERE id = %d", ID)

	_, err = repo.Db.Queryx(query)

	if err != nil {
		http.Error(w, "Não foi possivel deletar o registro", http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "registro deletado com sucesso")
	return
}
