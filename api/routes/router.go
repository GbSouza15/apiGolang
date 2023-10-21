package api

import (
	"net/http"

	api "github.com/GbSouza15/apiGolang/api/handlers"
)

func RoutesApi() {
	http.HandleFunc("/", api.HomeHandler)
}
