package main

import (
	"github.com/GbSouza15/apiGolang/api/routes"
	"github.com/GbSouza15/apiGolang/config/db"
)

func main() {

	data := db.InitDb()

	db.CreateSchemaAndTable(data)

	routes.RoutesApi(data)

	db.CloseConnection(data)
}
