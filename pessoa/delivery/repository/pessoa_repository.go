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

//Incluir inclui uma pessoa no banco de dados
func (rep *PessoaRepository) Incluir(pessoa model.Pessoa) (err error) {

	query := fmt.Sprintf("INSERT INTO pessoa (id,name,age,telephone) VALUES (%d,'%s',%d,'%s')", pessoa.Id, pessoa.Nome, pessoa.Idade, pessoa.Telefone)

	_, err = rep.Db.Exec(query)

	if err != nil {
		return err
	}

	return err
}

//Alterar altera um registro de pessoa
func (rep *PessoaRepository) Alterar(pessoa model.Pessoa) (err error) {

	query := fmt.Sprintf("UPDATE pessoa SET name = '%s', age = %d, telephone = '%s' WHERE id = %d", pessoa.Nome, pessoa.Idade, pessoa.Telefone, pessoa.Id)

	_, err = rep.Db.Exec(query)

	if err != nil {
		return err
	}

	return err
}

//Deletar deleta um registro de pessoa
func (rep *PessoaRepository) Deletar(ID int64) (err error) {
	query := fmt.Sprintf("DELETE FROM pessoa WHERE id = %d", ID)

	_, err = rep.Db.Exec(query)

	if err != nil {
		return err
	}

	return err
}
