package auth

import (
	"errors"
	"evermos/config"
	"evermos/helper"
	"evermos/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
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
	if err := config.DB.Where("notelp = ?", notelp).First(&user).Error; err != nil {
		return nil, errors.New("Nomor Telepon atau Password salah")
	}

	if !helper.CheckPassword(password, user.KataSandi) {
		return nil, errors.New("Nomor Telepon atau Password salah")
	}

	// Token expired dalam 7 hari
	claims := jwt.MapClaims{
		"id":       user.ID,
		"is_admin": user.IsAdmin,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(), //
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tkn.SignedString([]byte("secret_key"))
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{
		"nama":    user.Nama,
		"no_telp": user.Notelp,
		"email":   user.Email,
		"token":   tokenString,
	}

	return result, nil
}
