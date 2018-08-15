package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang-crud/model"
	"github.com/golang-crud/pessoa/delivery/service"
	"github.com/google/uuid"
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

var pessoaService = service.PessoaService{}

// ObterPessoas funcao que obtem pessoas
func ObterPessoas(w http.ResponseWriter, r *http.Request) {

	pessoas, err := pessoaService.ObterPessoas()

	if err != nil {
		http.Error(w, "Não encontramos nemhum item", http.StatusInternalServerError)
		fmt.Println("[main] [select pessoa] ", err.Error())
		return
	}

	json, _ := json.Marshal(pessoas)
	s := string(json[:])

	fmt.Fprint(w, s)

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

	id, err := uuid.NewUUID()

	if err != nil {
		http.Error(w, "Error ao ler json "+err.Error(), http.StatusBadRequest)
		return
	}

	pessoa.ID = id.String()

	err = pessoaService.CriarPessoa(pessoa)

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

	err = pessoaService.AlterarPessoa(pessoa)

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

	if pessoaID == "" {
		http.Error(w, "id vazio", http.StatusBadRequest)
		return
	}

	err := pessoaService.DeletarPessoa(pessoaID)

	if err != nil {
		http.Error(w, "Não foi possivel deletar o registro", http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "registro deletado com sucesso")
	return
}
