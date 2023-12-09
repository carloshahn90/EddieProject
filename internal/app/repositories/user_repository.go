package repositories

import (
	"github.com/carloshahn90/EddieProject/pkg/db"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) IsValidCredentials(username, password string) (bool, error) {
	// Consulta SQL para verificar as credenciais
	query := "SELECT COUNT(*) FROM usuario WHERE username = $1 AND password = $2"
	var count int
	err := db.DB.QueryRow(query, username, password).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
