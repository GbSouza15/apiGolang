package main

import (
	"fmt"
	"net/http"

	api "github.com/GbSouza15/apiGolang/api/routes"
	"github.com/GbSouza15/apiGolang/config/db"
)

func main() {

	err := db.InitDb()

	if err != nil {
		panic(err)
	}

	err = db.CreateSchemaAndTable()

	if err != nil {
		panic(err)
	}

	api.RoutesApi()

	fmt.Println("Server is running on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
