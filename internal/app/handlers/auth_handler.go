package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/carloshahn90/EddieProject/internal/app/services"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		authService: services.NewAuthService(),
	}
}

func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user map[string]string
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, "Erro ao decodificar as credenciais", http.StatusBadRequest)
		return
	}

	username, ok := user["username"]
	if !ok {
		http.Error(w, "Nome de usuário não fornecido", http.StatusBadRequest)
		return
	}

	password, ok := user["password"]
	if !ok {
		http.Error(w, "Senha não fornecida", http.StatusBadRequest)
		return
	}

	// Chamar o serviço para autenticar
	token, err := ah.authService.Authenticate(username, password)
	if err != nil {
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}

	// Retornar o token como resposta
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"token": "` + token + `"}`))
}
