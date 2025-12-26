package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// AuthMiddleware untuk validasi user umum
func AuthMiddleware(c *fiber.Ctx) error {
	// Mengambil token dari header 'token'
	authHeader := c.Get("token")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  false,
			"message": "Unauthorized",
			"errors":  []string{"Token is required"},
		})
	}

	// Cek validasi token
	token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret_key"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Locals("user_id", claims["id"])
		c.Locals("is_admin", claims["is_admin"])
		return c.Next()
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"status":  false,
		"message": "Unauthorized",
		"errors":  []string{err.Error()},
	})
}

// AdminMiddleware untuk proteksi rute khusus admin
func AdminMiddleware(c *fiber.Ctx) error {
	isAdmin, ok := c.Locals("is_admin").(bool)

	if !ok || !isAdmin {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  false,
			"message": "Forbidden: Hanya Admin yang dapat mengakses ini",
		})
	}

	return c.Next()
}
