package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDb() error {
	var err error

	godotenv.Load()

	var (
		user    = os.Getenv("USER")
		pass    = os.Getenv("PASS")
		dbname  = os.Getenv("DB")
		host    = os.Getenv("HOST")
		port    = os.Getenv("PORT")
		sslmode = os.Getenv("SSLMODE")
	)

	connectString := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=%s", user, dbname, pass, host, port, sslmode)

	db, err = sql.Open("postgres", connectString)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database successfully.")

	return nil
}

func CreateSchemaAndTable() error {

	createSchemaSQL := "CREATE SCHEMA IF NOT EXISTS api"

	_, err := db.Exec(createSchemaSQL)

	if err != nil {
		return err
	}

	createTableSQL := "CREATE TABLE IF NOT EXISTS api.teste (id SERIAL PRIMARY KEY, name VARCHAR(255) NOT NULL)"

	_, err = db.Exec(createTableSQL)

	if err != nil {
		return err
	}

	checkExistSchema := "SELECT EXISTS(SELECT 1 FROM information_schema.tables WHERE table_schema = 'api')"

	var exists bool

	err = db.QueryRow(checkExistSchema).Scan(&exists)

	if err != nil {
		return err
	}

	if exists {
		fmt.Println("Schema já existe!")

		return nil
	}

	fmt.Println("Schema e tabela criados com sucesso!")
	return nil
}
