package repositories

import (
	"fmt"

	"github.com/carloshahn90/EddieProject/pkg/db"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) IsValidCredentials(username, password string) (int, error) {
	fmt.Printf("username: %s, password: %s \n", username, password)
	// Consulta SQL para verificar as credenciais
	query := "SELECT id FROM \"user\" WHERE user_name = $1 AND user_password = $2"
	var id int
	err := db.DB.QueryRow(query, username, password).Scan(&id)
	if err != nil {
		fmt.Printf("ERROR: %v \n", err)
		return 0, err
	}
	fmt.Printf("id: %v \n", id)
	return id, nil
}
