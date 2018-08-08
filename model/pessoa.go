package model

//Pessoa representa uma pessoa
type Pessoa struct {
	ID       string `json:"id" db:"id"`
	Nome     string `json:"nome" db:"name"`
	Idade    int    `json:"idade" db:"age"`
	Telefone string `json:"telefone" db:"telephone"`
}
