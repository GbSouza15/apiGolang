package main

import (
	"fmt"
	"net/http"

	api "github.com/GbSouza15/apiGolang/api/routes"
)

func main() {

	api.RoutesApi()

	fmt.Println("Server is running on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
