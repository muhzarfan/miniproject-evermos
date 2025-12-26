package wilayah

import (
	"github.com/gofiber/fiber/v2"
)

// Mengambil daftar provinsi
func GetProvincesHandler(c *fiber.Ctx) error {
	res, err := GetProvinces()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengambil data provinsi",
		})
	}
	return c.JSON(res)
}

// Mengambil daftar kota berdasarkan id_provinsi
func GetRegenciesHandler(c *fiber.Ctx) error {
	idProvinsi := c.Params("id_provinsi")
	res, err := GetRegencies(idProvinsi)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengambil data kota",
		})
	}
	return c.JSON(res)
}
