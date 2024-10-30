package models

import "gorm.io/gorm"

// User representa o modelo de um usuário.
type User struct {
    gorm.Model
	ID 		 uint64 `gorm:"primaryKey;autoincrement"`
    Username string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
}
