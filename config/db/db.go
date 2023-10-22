package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDb() *sql.DB {
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

	db, err := sql.Open("postgres", connectString)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database successfully.")

	return db
}

func CloseConnection(db *sql.DB) {
	defer db.Close()
}

func CreateSchemaAndTable(db *sql.DB) {

	createSchemaSQL := "CREATE SCHEMA IF NOT EXISTS api"

	_, err := db.Exec(createSchemaSQL)

	if err != nil {
		fmt.Println("Error creating schema: ", err.Error())
		return
	}

	createTableSQL := "CREATE TABLE IF NOT EXISTS api.teste (id SERIAL PRIMARY KEY, name VARCHAR(255) NOT NULL)"

	_, err = db.Exec(createTableSQL)

	if err != nil {
		fmt.Println("Error creating table: ", err.Error())
		return
	}

	checkExistSchema := "SELECT EXISTS(SELECT 1 FROM information_schema.tables WHERE table_schema = 'api')"

	var exists bool

	err = db.QueryRow(checkExistSchema).Scan(&exists)

	if err != nil {
		fmt.Println("Error checking schema: ", err.Error())
		return
	}

	if exists {
		fmt.Println("Schema exists")
		return
	}

	fmt.Println("Schema e tabela criados com sucesso!")
}
