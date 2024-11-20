package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/patrick/user-module-go/internal/database"
	"github.com/patrick/user-module-go/internal/handlers"
	"github.com/patrick/user-module-go/internal/middleware"
)

func main() {
	app := fiber.New()

	// Conectar ao banco de dados
	_, err := database.ConnectToDatabase()
	if err != nil {
		fmt.Println("Não foi possível conectar ao banco de dados: %v", err)
	}

	// Rotas públicas
	app.Post("/register", handlers.RegisterUser)
	app.Post("/login", handlers.Login)

	// Rota protegida
	app.Get("/userInfos", middleware.Protect(), handlers.GetUserInfos)

	// Obtenha a porta da variável de ambiente ou defina uma padrão
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Usando a porta 3000 como fallback se a variável não estiver definida
	}

	// Iniciar o servidor na porta fornecida e garantir que ele ouça em todas as interfaces (0.0.0.0)
	if err := app.Listen("0.0.0.0:" + port); err != nil {
		panic(err)
	}
}
