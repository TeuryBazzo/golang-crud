package repo

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	//github.com/lib/pq drive para o postgre
	_ "github.com/lib/pq"
)

//Db é o manipulador do banco
var Db *sqlx.DB

//AbrirConexao abrir conexao com o banco de dados
func AbrirConexao() (err error) {
	connStr := "user=postgres password=admin dbname=golangdb port=5432 sslmode=disable"

	Db, err = sqlx.Open("postgres", connStr)

	if err != nil {
		fmt.Println("[repo] Erro ao abrir conexao com o banco de dados ", err.Error())
		return
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println("[repo] funcao que abre a conexao com banco de dados", err.Error())
		return
	}

	return
}
