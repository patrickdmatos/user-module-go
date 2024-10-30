package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase estabelece uma conexão com o banco de dados SQLite.
func ConnectDatabase() {
    var err error
    DB, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }
}

// Migrate realiza a migração dos modelos para o banco de dados.
func Migrate() {
    DB.AutoMigrate(&models.User{})
}
