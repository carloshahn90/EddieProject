package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/carloshahn90/EddieProject/internal/app/services"
	"github.com/go-chi/chi"
)

func (ah *Handler) GetDog(w http.ResponseWriter, r *http.Request) {

	userId := chi.URLParam(r, "userId")
	tokenString := r.Header.Get("Authorization")

	id, err := services.ValidToken(tokenString, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	dogs := ah.dogService.GetDog(id)

	responseJSON, _ := json.Marshal(dogs)
	/*response := map[string]string{
		"userId": fmt.Sprint(id),
		"dog":    "golden retriever",
		"token":  tokenString,
	}*/

	// Retornar a resposta em formato JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
