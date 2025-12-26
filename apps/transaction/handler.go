package transaction

import (
	"evermos/models"

	"github.com/gofiber/fiber/v2"
)

// Membuat Transaksi (POST)
func CreateTrxHandler(c *fiber.Ctx) error {
	// Ambil id user
	userID := c.Locals("user_id").(float64)

	var input models.Trx

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid input",
		})
	}

	res, err := CreateTransaction(uint(userID), input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  true,
		"message": "Berhasil membuat invoice",
		"data":    res,
	})
}

// Mengambil Data Transaksi (GET)
func GetAllTrxHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(float64)
	namaProduk := c.Query("nama_produk")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	res, err := GetAllTransaction(uint(userID), namaProduk, startDate, endDate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal mengambil riwayat",
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Berhasil mengambil data",
		"data":    res,
	})
}
