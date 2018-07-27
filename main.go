package main

import (
	"fmt"
	"net/http"
	"projects/golang-crud/repo"
)

func main() {

	err := repo.AbrirConexao()

	if err != nil {
		fmt.Println("[main] [abrir conexao] ", err.Error())
		return
	}

	http.HandleFunc("/pessoa", NewPessoaHTTPHandler)

	http.ListenAndServe(":9000", nil)
}

//NewPessoaHTTPHandler redireciona o request do endpoint /pessoa
func NewPessoaHTTPHandler(w http.ResponseWriter, r *http.Request) {

}
