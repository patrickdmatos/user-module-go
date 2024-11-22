package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/patrickdmatos/user-module-go/internal/services"
)

// Função de handler para registrar o usuário
func RegisterUser(c *fiber.Ctx) error {
	var body struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Verifica se a requisição está correta
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	// Tenta registrar o usuário
	err := services.RegisterUser(body.Username, body.Name, body.Email, body.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email already registered, please try to login",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

// Função de handler para login e gerar o token JWT
func Login(c *fiber.Ctx) error {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Parseia o corpo da requisição
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	// Autentica o usuário e gera o token
	token, err := services.AutenticateUser(body.Email, body.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "User authenticated successfully",
		"tokenSession": token,
	})
}

// Função para obter as informações do usuário (somente para rotas protegidas)
func GetUserInfos(c *fiber.Ctx) error {
   	// Pegue o email do usuário a partir do contexto com type assertion
	email, ok := c.Locals("email").(string)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email not found in context",
		})
	}

	// Chama o serviço para buscar o usuário no banco
	user, err := services.GetUserByEmail(email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Retorna os dados do usuário
	return c.JSON(fiber.Map{
		"name":     user.Name,
		"email":    user.Email,
	})
}

