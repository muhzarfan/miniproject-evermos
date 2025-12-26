package produk

import (
	"evermos/config"
	"evermos/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Membuat produk baru (POST)
func CreateProdukHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(float64)

	// Cari ID Toko berdasarkan User ID
	var toko models.Toko
	config.DB.Where("id_user = ?", userID).First(&toko)

	var input models.Produk
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Format data tidak valid",
		})
	}

	input.IDToko = toko.ID

	res, err := CreateProduk(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal buat produk",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  true,
		"message": "Berhasil membuat data",
		"data":    res.ID,
	})
}

// Mengambil data produk dengan Filter (GET)
func GetAllProdukHandler(c *fiber.Ctx) error {
	nama := c.Query("nama")
	categoryID := c.Query("id_category")

	res, err := GetAllProduk(nama, categoryID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal mengambil data",
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Berhasil mengambil semua data",
		"data":    res,
	})
}

// Mengambil produk berdasarkan ID (GET)
func GetProdukByIDHandler(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	res, err := GetProdukByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Produk tidak ditemukan",
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Berhasil mengambil data berdasarkan ID",
		"data":    res,
	})
}

// Memperbarui produk (PUT)
func UpdateProdukHandler(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	userID := c.Locals("user_id").(float64)

	var toko models.Toko
	config.DB.Where("id_user = ?", userID).First(&toko)

	var input models.Produk
	c.BodyParser(&input)

	err := UpdateProduk(uint(id), toko.ID, input)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Berhasil memperbarui data",
	})
}

// Menghapus data produk (DELETE)
func DeleteProdukHandler(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	userID := c.Locals("user_id").(float64)

	var toko models.Toko
	config.DB.Where("id_user = ?", userID).First(&toko)

	err := DeleteProduk(uint(id), toko.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal hapus produk",
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Berhasil menghapus data",
	})
}
