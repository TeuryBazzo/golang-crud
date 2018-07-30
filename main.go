package main

import (
	"fmt"
	"net/http"

	"github.com/golang-crud/article/delivery/handler"
	"github.com/golang-crud/repo"
)

func main() {

	err := repo.AbrirConexao()

	if err != nil {
		fmt.Println("[main] [abrir conexao] ", err.Error())
		return
	}

	http.HandleFunc("/pessoa", handler.NewPessoaHTTPHandler)

	http.ListenAndServe(":9000", nil)
}
