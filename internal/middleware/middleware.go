package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/patrick/user-module-go/internal/services"
)

var jwtKey = []byte("secrect_key") // Mesma chave secreta usada para gerar o token

func Protect() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extrai o token do cabeçalho Authorization
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).SendString("Token não fornecido")
		}

		// O token está no formato "Bearer token", então extraímos o valor
		tokenString := strings.Split(authHeader, " ")[1]

		// Verifica e valida o token
		token, err := jwt.ParseWithClaims(tokenString, &services.Claims{}, func(token *jwt.Token) (interface{}, error) {
			// Verifica se o token foi assinado com o algoritmo correto
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("não foi possível verificar a assinatura do token")
			}
			return jwtKey, nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString(fmt.Sprintf("Token inválido: %v", err))
		}

		// Extraímos os claims do token
		claims, ok := token.Claims.(*services.Claims)
		if !ok || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).SendString("Token inválido")
		}

		// Atribui os dados do usuário no contexto (opcional)
		c.Locals("email", claims.Email)

		return c.Next()
	}
}
