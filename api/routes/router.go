package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	api "github.com/GbSouza15/apiGolang/api/handlers"
)

func RoutesApi(db *sql.DB) {
	h := api.New(db)

	http.HandleFunc("/articles", h.GetAllArticles)

	fmt.Println("Server is running on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
