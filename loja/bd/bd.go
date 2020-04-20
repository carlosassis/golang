package bd

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaBD() *sql.DB {
	conexao := "user=postgres dbname=storedb port=54320 password=mysecretpassword host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}
