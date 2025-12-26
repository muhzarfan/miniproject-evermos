package auth

import (
	"evermos/models"

	"github.com/gofiber/fiber/v2"
)

// REGISTER
func RegisterHandler(c *fiber.Ctx) error {
	var input models.User

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid input",
		})
	}

	user, err := Register(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal Register Pengguna",
			"errors":  []string{err.Error()},
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Berhasil membuat data",
		"data":    user,
	})
}

// LOGIN
func LoginHandler(c *fiber.Ctx) error {
	var input struct {
		NoTelp    string `json:"notelp"`
		KataSandi string `json:"kata_sandi"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid input",
		})
	}

	result, err := Login(input.NoTelp, input.KataSandi)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal Login",
			"errors":  []string{err.Error()},
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Berhasil Login",
		"data":    result,
	})
}

// LOGOUT
func LogoutHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Berhasil Logout",
		"data":    nil,
	})
}
