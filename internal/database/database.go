package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() (*gorm.DB, error) {
	// Carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	// Obtém as variáveis de ambiente
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DATABASE")
	port := os.Getenv("POSTGRES_PORT")

	// Verifique se as variáveis estão vazias
	if user == "" || password == "" || port == "" {
		return nil, fmt.Errorf("env variables POSTGRES_USER, POSTGRES_PASSWORD, and POSTGRES_PORT must be set")
	}

	// Cria o DSN com variáveis de ambiente
	dsn := fmt.Sprintf("postgresql://%s:%s@%s.oregon-postgres.render.com/%s", user, password, host, dbname)

	// Conecta ao banco de dados
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	DB = db
	return db, nil
}
