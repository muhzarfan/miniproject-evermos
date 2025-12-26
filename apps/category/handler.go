package category

import (
	"evermos/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Membuat kategori baru (POST)
func CreateCategoryHandler(c *fiber.Ctx) error {
	var input models.Category

	// Fiber menggunakan BodyParser untuk JSON
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid input",
		})
	}

	res, err := CreateCategory(input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal membuat kategori",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  true,
		"message": "Berhasil membuat data",
		"data":    res.ID,
	})
}

// Mengambil semua kategori (GET)
func GetAllCategoryHandler(c *fiber.Ctx) error {
	res, err := GetAllCategories()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal mengambil data",
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Berhasil mendapatkan data",
		"data":    res,
	})
}

// Memperbarui kategori (PUT)
func UpdateCategoryHandler(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var input models.Category
	c.BodyParser(&input)

	err := UpdateCategory(uint(id), input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal update kategori",
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Berhasil memperbarui data",
		"data":    "",
	})
}

// Menghapus kategori (DELETE)
func DeleteCategoryHandler(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	err := DeleteCategory(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal menghapus kategori",
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Berhasil menghapus data",
		"data":    "",
	})
}
