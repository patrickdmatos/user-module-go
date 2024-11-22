package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/patrickdmatos/user-module-go/internal/database"
	"github.com/patrickdmatos/user-module-go/internal/models"
	"golang.org/x/crypto/bcrypt"
)


var jwtKey = []byte("secrect_key") // Chave secreta para assinar o token JWT

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func AutenticateUser(email, password string) (string, error) {
	var user models.User

	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
			if err.Error() == "record not found" {
			return "", fmt.Errorf("usuário não encontrado")
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid password")
	}

	//Gerar token JWT
	expirationTime := time.Now().Add(24 * time.Hour) // Token expira em 24h
	claims := &Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // Converte para Unix timestamp
			Issuer:    "user-module-go",
		},
	}

	//Cria o token JWT com os claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Assina o token com a chave secreta
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("não foi possível gerar o token: %v", err)
	}

	return tokenString, nil
}