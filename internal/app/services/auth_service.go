package services

import (
	"github.com/carloshahn90/EddieProject/internal/app/repositories"
)

type AuthService struct {
	userRepo repositories.UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		userRepo: repositories.NewUserRepository(),
	}
}

func (as *AuthService) Authenticate(username, password string) (string, error) {
	// Verificar as credenciais no repositório
	isValid, err := as.userRepo.IsValidCredentials(username, password)
	if err != nil {
		return "", err
	}

	if isValid {
		// Gerar token
		return GenerateToken(username)
	}

	return "", nil
}
