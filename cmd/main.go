package main

import (
	"user-module/internal/database"
	"user-module/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    database.ConnectDatabase()
    database.Migrate() // Realiza a migração do modelo

    app.Post("/api/users/register", handlers.RegisterUser)
    // Adicione mais rotas conforme necessário

    app.Listen(":3000")
}
