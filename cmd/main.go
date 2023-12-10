package main

import (
	"fmt"
	"net/http"

	"github.com/carloshahn90/EddieProject/internal/app/handlers"
	"github.com/carloshahn90/EddieProject/pkg/db"
	"github.com/go-chi/chi"
)

func main() {
	// Iniciar conexão com o banco de dados PostgreSQL
	db.InitDB("localhost", "5432", "postgres", "root", "EddieProject")

	// Iniciar roteador e manipuladores
	authHandler := handlers.NewAuthHandler()

	// Configurar rotas usando chi.NewRouter()
	r := chi.NewRouter()
	authHandler.SetupRoutes(r)

	http.ListenAndServe(":8080", r)
	fmt.Println("Servidor no ar na porta 8080")
}
