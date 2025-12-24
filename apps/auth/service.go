package auth

import (
	"errors"
	"evermos/config"
	"evermos/helper"
	"evermos/models"
)

// REGISTER
func Register(input models.User) (models.User, error) {
	var existingUser models.User
	// Cek ketersediaan email atau no telepon
	if err := config.DB.Where("email = ? OR notelp = ?", input.Email, input.Notelp).First(&existingUser).Error; err == nil {
		return input, errors.New("Email atau Nomor Telepon sudah terdaftar")
	}

	// Hashing password
	hashedPwd, _ := helper.HashPassword(input.KataSandi)
	input.KataSandi = hashedPwd

	// Menyimpan user ke database
	if err := config.DB.Create(&input).Error; err != nil {
		return input, err
	}

	// Membuat toko otomatis setelah daftar user
	newToko := models.Toko{
		IDUser:   input.ID,
		NamaToko: "Toko " + input.Nama,
		UrlFoto:  "",
	}
	config.DB.Create(&newToko)

	return input, nil
}

// LOGIN
func Login(notelp string, password string) (map[string]interface{}, error) {
	var user models.User
	// Cari user berdasarkan notelp
	if err := config.DB.Where("notelp = ?", notelp).First(&user).Error; err != nil {
		return nil, errors.New("Nomor Telepon atau Password salah")
	}

	// verifikasi password menggunakan helper bcrypt
	if !helper.CheckPassword(password, user.KataSandi) {
		return nil, errors.New("Nomor Telepon atau Password salah")
	}

	// jwt token
	token, err := helper.GenerateToken(user.ID, user.IsAdmin)
	if err != nil {
		return nil, err
	}

	// menyusun respons data
	result := map[string]interface{}{
		"nama":    user.Nama,
		"no_telp": user.Notelp,
		"email":   user.Email,
		"token":   token,
	}

	return result, nil
}
