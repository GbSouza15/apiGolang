package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Server is running on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
