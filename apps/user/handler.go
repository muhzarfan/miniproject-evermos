package user

import (
	"evermos/apps/wilayah"
	"evermos/models"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Struct untuk menampilkan nama provinsi dan kota saat ambil data user
type UserProfileResponse struct {
	models.User
	NamaProvinsi string `json:"nama_provinsi"`
	NamaKota     string `json:"nama_kota"`
}

// Mengambil data user yang login (GET)
func GetMyProfileHandler(c *fiber.Ctx) error {
	// Ambil id user
	userID := c.Locals("user_id").(float64)

	userData, err := GetMyProfile(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "User tidak ditemukan",
		})
	}

	response := UserProfileResponse{User: userData}

	// Integrasi data API Wilayah untuk provinsi dan kota
	if userData.IDProvinsi != "" {
		provList, _ := wilayah.GetProvinces()
		for _, p := range provList {
			if fmt.Sprintf("%v", p["id"]) == userData.IDProvinsi {
				response.NamaProvinsi = p["name"].(string)
				break
			}
		}
	}

	if userData.IDKota != "" {
		kotaList, _ := wilayah.GetRegencies(userData.IDProvinsi)
		for _, k := range kotaList {
			if fmt.Sprintf("%v", k["id"]) == userData.IDKota {
				response.NamaKota = k["name"].(string)
				break
			}
		}
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Berhasil mengambil data",
		"data":    response,
	})
}

// Perbarui Akun User yang Login (PUT)
func UpdateProfileHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(float64)

	var input models.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid input",
		})
	}

	err := UpdateProfile(uint(userID), input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal update profil",
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Update profil berhasil",
	})
}

// Ubah Password User (PUT)
func ChangePasswordHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(float64)

	var input struct {
		PasswordBaru string `json:"password_baru"`
	}
	c.BodyParser(&input)

	if len(input.PasswordBaru) < 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Password minimal 6 karakter",
		})
	}

	err := ChangePassword(uint(userID), input.PasswordBaru)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal mengganti password",
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Password berhasil diganti",
	})
}

// Membuat alamat baru (POST)
func CreateAlamatHandler(c *fiber.Ctx) error {
	var input models.Alamat
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid input",
		})
	}

	userIDFloat := c.Locals("user_id").(float64)
	input.IDUser = uint(userIDFloat)

	res, err := CreateAlamat(input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal menambah alamat",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  true,
		"message": "Berhasil menyimpan data",
		"data":    res.ID,
	})
}

// Mengambil data alamat milik user (GET)
func GetMyAlamatHandler(c *fiber.Ctx) error {
	userIDFloat := c.Locals("user_id").(float64)

	res, err := GetMyAlamat(uint(userIDFloat))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal mengambil data alamat",
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Berhasil mendapat data",
		"data":    res,
	})
}

// Memperbarui alamat (PUT)
func UpdateAlamatHandler(c *fiber.Ctx) error {
	idStr := c.Params("id")
	userID := c.Locals("user_id").(float64)

	var input models.Alamat
	c.BodyParser(&input)

	addrID, _ := strconv.ParseUint(idStr, 10, 32)

	err := UpdateAlamat(uint(addrID), uint(userID), input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal update atau data tidak ditemukan",
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Berhasil update data.",
		"data":    "",
	})
}

// Menghapus alamat (DELETE)
func DeleteAlamatHandler(c *fiber.Ctx) error {
	idStr := c.Params("id")
	userID := c.Locals("user_id").(float64)

	addrID, _ := strconv.ParseUint(idStr, 10, 32)

	err := DeleteAlamat(uint(addrID), uint(userID))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal menghapus data",
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Berhasil menghapus data.",
		"data":    "",
	})
}
