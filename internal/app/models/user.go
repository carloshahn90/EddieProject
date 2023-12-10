package models

// Adapte conforme necessário com a estrutura real do seu usuário
type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
