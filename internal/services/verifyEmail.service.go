package services

import (
	"fmt"

	"github.com/patrickdmatos/api-all-in-go/database"
	"github.com/patrickdmatos/api-all-in-go/models"
)

// Verifica se o email já está cadastrado no banco de dados
func IsEmailTaken(email string) (bool, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	fmt.Println(err)
	if err != nil {
		if err.Error() == "record not found" {
			return false, nil // Não encontrou o usuário, então o email não está registrado
		}
		return false, err // Erro inesperado ao consultar o banco de dados
	}
	return true, nil // Email já registrado
}