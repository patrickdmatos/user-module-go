package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/patrickdmatos/api-shared-library-go/middleware"
)

var jwtKey = []byte("secret_key") // Mesma chave secreta usada para gerar o token

func Protect() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extrai o token do cabeçalho Authorization
		authHeader := c.Get("Authorization")
		claims := middleware.Protect(authHeader)

		return c.JSON(claims)
	}
}
