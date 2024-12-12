package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickdmatos/api-shared-library-go/database"
	"github.com/patrickdmatos/user-module-go/internal/handlers"
	"github.com/patrickdmatos/user-module-go/middleware"
)

func main() {
	app := fiber.New()

	// Conectar ao banco de dados
	_, err := database.ConnectToDatabase()
	if err != nil {
		panic(err)
	}

	// Rotas públicas
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("Hello World, lets GO!")
	})
	app.Post("/register", handlers.RegisterUser)
	app.Post("/login", handlers.Login)

	// Rota protegida
	app.Get("/userInfos", middleware.Protect(), handlers.GetUserInfos)

	// Obtenha a porta da variável de ambiente ou defina uma padrão
	 port := os.Getenv("PORT")
    if port == "" {
        fmt.Println("Erro: A variável PORT não foi configurada!")
        os.Exit(1)
    }

    fmt.Printf("A API está rodando na porta: %s\n", port)

	// Iniciar o servidor na porta fornecida e garantir que ele ouça em todas as interfaces (0.0.0.0)
	if err := app.Listen("0.0.0.0:" + port); err != nil {
		panic(err)
	}
}
