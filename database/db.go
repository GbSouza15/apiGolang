package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConectDb() {
	db, err := sql.Open("postgres", "user=postgres dbname=api password=gabriel2004 host=localhost port=5432 sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Conectado com sucesso!")
}
