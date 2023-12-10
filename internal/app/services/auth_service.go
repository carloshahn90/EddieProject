package services

import (
	"errors"
	"strconv"
	"time"

	"github.com/carloshahn90/EddieProject/internal/app/repositories"
	"github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

var (
	secretKey = []byte("TESTE")
)

// Claims representa os detalhes do token
type Claims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func NewAuthService() *AuthService {
	return &AuthService{
		userRepo: repositories.NewUserRepository(),
	}
}

func (as *AuthService) Authenticate(username, password string) (string, error) {
	// Verificar as credenciais no repositório
	id, err := as.userRepo.IsValidCredentials(username, password)
	if err != nil {
		return "", err
	}
	// Gerar token
	return generateToken(id, username)
}

func generateToken(id int, username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token expira em 1 dia

	claims := &Claims{
		Id:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)

	return tokenString, err
}

func ValidToken(tokenString string, userId string) (int, error) {

	userIdInt, _ := strconv.Atoi(userId)
	if tokenString == "" {
		//http.Error(w, "Token de autenticação ausente", http.StatusUnauthorized)
		return 0, errors.New("token de autenticação ausente")
	}

	// Parse do token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return 0, errors.New("erro ao parsear o token")
	}

	// Verificação do token
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {

		if userIdInt == claims.Id {
			return claims.Id, nil
		}
	}
	return 0, errors.New("token inválido")
}
