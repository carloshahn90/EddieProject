package models

// Adapte conforme necessário com a estrutura real do seu usuário
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
