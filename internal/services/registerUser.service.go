package services

import (
	"fmt"

	"github.com/patrickdmatos/api-all-in-go/database"
	"github.com/patrickdmatos/api-all-in-go/models"
	"golang.org/x/crypto/bcrypt"
)

//Registra um novo usuario na base
func RegisterUser(username, name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	emailAlreadyUsed, err := IsEmailTaken(email)

	if err != nil {
		return err
	}

	if emailAlreadyUsed == true {
		return fmt.Errorf("o email já está cadastrado")
	}

	user := models.User{
		Name:     name,
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	return database.DB.Create(&user).Error
}
