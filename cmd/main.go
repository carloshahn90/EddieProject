package main

import (
	"net/http"

	"github.com/carloshahn90/EddieProject/internal/app/handlers"
	"github.com/carloshahn90/EddieProject/pkg/db"
)

func main() {
	// Iniciar conexão com o banco de dados PostgreSQL
	db.InitDB("seu_host", "sua_porta", "seu_usuario", "sua_senha", "seu_nome_db")

	// Iniciar roteador e manipuladores
	authHandler := handlers.NewAuthHandler()

	http.HandleFunc("/login", authHandler.Login)

	// Iniciar o servidor
	http.ListenAndServe(":8080", nil)
}
