package toko

import (
	"evermos/models"
	"io"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Mengambil data toko user yang login (GET)
func GetMyTokoHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(float64)

	res, err := GetMyToko(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Toko tidak ditemukan",
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Berhasil mendapat data.",
		"data":    res,
	})
}

// Memperbarui profil toko dan unggah foto (PUT)
func UpdateTokoHandler(c *fiber.Ctx) error {
	tokoIDStr := c.Params("id_toko")
	userID := c.Locals("user_id").(float64)

	id, err := strconv.ParseUint(tokoIDStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "ID Toko tidak valid",
		})
	}

	// Ambil file foto dari form-data
	fileHeader, err := c.FormFile("url_foto")
	var fileName string
	if err == nil {
		// Memastikan folder uploads ada
		os.MkdirAll("uploads", os.ModePerm)
		fileName = "uploads/" + fileHeader.Filename

		// Upload file
		src, err := fileHeader.Open()
		if err == nil {
			defer src.Close()
			dst, err := os.Create(fileName)
			if err == nil {
				defer dst.Close()
				io.Copy(dst, src)
			}
		}
	}

	// Ambil nilai text dari form-data
	namaToko := c.FormValue("nama_toko")

	input := models.Toko{
		NamaToko: namaToko,
		UrlFoto:  fileName,
	}

	err = UpdateToko(uint(id), uint(userID), input)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal update toko atau akses dilarang",
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Berhasil update toko.",
		"data":    "",
	})
}
