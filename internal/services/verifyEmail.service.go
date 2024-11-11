package services

import (
	"github.com/patrick/user-module-go/internal/database"
	"github.com/patrick/user-module-go/internal/models"
)

// Verifica se o email já está cadastrado no banco de dados
func IsEmailTaken(email string) (bool, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err.Error() == "record not found" {
			return false, nil // Não encontrou o usuário, então o email não está registrado
		}
		return false, err // Erro inesperado ao consultar o banco de dados
	}
	return true, nil // Email já registrado
}