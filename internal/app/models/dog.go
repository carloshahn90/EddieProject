package models

// Adapte conforme necessário com a estrutura real do seu usuário
type Dog struct {
	Id     int64  `json:"id"`
	UserId int64  `json:"user_id"`
	Breed  string `json:"breed"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
}
