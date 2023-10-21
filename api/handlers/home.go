package api

import (
	"encoding/json"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello World"})

	} else {
		http.Error(w, "Erro na requisição", http.StatusBadRequest)
	}
}
