package repository

import (
	"fmt"

	"github.com/golang-crud/model"
	"github.com/jmoiron/sqlx"
)

//PessoaRepository representa a classe de acesso a base de dados
type PessoaRepository struct {
	Db *sqlx.DB
}

//ObterTodos obtem todas as pessoas
func (rep *PessoaRepository) ObterTodos() (pessoas []model.Pessoa, err error) {

	pessoas = []model.Pessoa{}

	rows, err := rep.Db.Queryx("select * from Pessoa")

	if err != nil {
		return nil, err
	}

	for rows.Next() {

		pessoa := model.Pessoa{}

		err = rows.StructScan(&pessoa)

		pessoas = append(pessoas, pessoa)

		if err != nil {
			fmt.Println("[main] [select pessoa] ", err.Error())
			return nil, err
		}
	}

	return pessoas, nil
}
