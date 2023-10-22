package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GbSouza15/apiGolang/api/models"
)

func (h handler) GetAllArticles(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}

	fmt.Println("GetAllArticles")

	results, err := h.DB.Query("SELECT * FROM api.teste")

	if err != nil {
		fmt.Println("Error getting articles: ", err.Error())
		return
	}

	var articles = make([]models.Articles, 0)

	for results.Next() {
		var article models.Articles

		err = results.Scan(&article.Id, &article.Name)

		if err != nil {
			fmt.Println("Error scanning articles: ", err.Error())
			return
		}

		articles = append(articles, article)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(articles)
}
