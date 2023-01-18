package db

import (
	"database/sql"

	_ "github.com/lib/pq" // O "_" significa que vai utilizar essa biblioteca apenas durante o tempo de execução da aplicação.
)

func ConnectDB() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=0506201Pb host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}
	return db
}
