package api

import (
	"encoding/json"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello World"})

		// está codificando um mapa (map) em formato JSON e escrevendo o resultado no objeto w, que é uma instância de http.ResponseWriter.

		// Aqui está o que está acontecendo em detalhes:

		// json.NewEncoder(w) cria um novo codificador JSON que está configurado para escrever os dados no objeto w.

		// .Encode(map[string]string{"message": "Hello World"}) é usado para codificar o mapa em formato JSON. Neste caso, o mapa contém uma única chave "message" com o valor "Hello World".

		// O resultado da codificação JSON é então escrito no http.ResponseWriter w, que é usado para enviar uma resposta HTTP de volta ao cliente que fez a solicitação.

		// Em resumo, essa linha de código está retornando uma resposta JSON contendo o objeto {"message": "Hello World"} para solicitações HTTP GET. Se a solicitação não for do tipo GET, ele retornará um erro HTTP 400 (Bad Request) para indicar que a solicitação não foi processada devido a um erro na requisição.

	} else {
		http.Error(w, "Erro na requisição", http.StatusBadRequest)
	}
}
