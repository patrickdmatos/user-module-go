package services

import (
	"github.com/patrick/user-module-go/internal/database"
	"github.com/patrick/user-module-go/internal/models"
)

func GetUserByEmail(email string) (*models.User, error) {
    var user models.User
    result := database.DB.Where("email = ?", email).First(&user)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}