package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/patrick/user-module-go/internal/database"
	"github.com/patrick/user-module-go/internal/handlers"
	"github.com/patrick/user-module-go/internal/middleware"
	"github.com/patrick/user-module-go/internal/models"
)

func main() {
	app := fiber.New()

	// Conectar ao banco de dados
	database.ConnectDatabase()
	database.DB.AutoMigrate(&models.User{})

	// Rotas públicas
	app.Post("/register", handlers.RegisterUser)
	app.Post("/login", handlers.Login)

	// Rota protegida
	app.Get("/userInfos", middleware.Protect(), handlers.GetUserInfos)

	// Iniciar o servidor
	app.Listen(":3000")
}
