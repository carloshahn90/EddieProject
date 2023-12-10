package handlers

import (
	"net/http"

	"github.com/carloshahn90/EddieProject/internal/app/services"
	"github.com/go-chi/chi"
)

type Handler struct {
	authService *services.AuthService
	dogService  *services.DogService
}

func NewAuthHandler() *Handler {
	return &Handler{
		authService: services.NewAuthService(),
		dogService:  services.NewDogService(),
	}
}

func (ah *Handler) SetupRoutes(r chi.Router) {
	r.Post("/login", ah.Login)
	r.Route("/user/{userId}", func(r chi.Router) {
		r.Use(ah.AuthMiddleware) // Middleware para verificar o token em rotas específicas
		r.Get("/get-dog", ah.GetDog)
	})
}

// Adicione esta função para o middleware de autenticação
func (ah *Handler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Token de autenticação ausente", http.StatusUnauthorized)
			return
		}

		// Aqui você pode adicionar a lógica para validar o token usando seu serviço de autenticação
		// por exemplo, chamando ah.authService.ValidateToken(token)

		// Chame o próximo manipulador se o token for válido
		next.ServeHTTP(w, r)
	})
}
