package user

import (
	"evermos/config"
	"evermos/models"

	"golang.org/x/crypto/bcrypt"
)

// USER

// Ambil data profil user yang login
func GetMyProfile(userID uint) (models.User, error) {
	var user models.User
	err := config.DB.Preload("Alamat").First(&user, userID).Error
	return user, err
}

// Update data profil
func UpdateProfile(userID uint, input models.User) error {
	return config.DB.Model(&models.User{}).Where("id = ?", userID).Updates(input).Error
}

// Ganti Password
func ChangePassword(userID uint, passwordBaru string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordBaru), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return config.DB.Model(&models.User{}).Where("id = ?", userID).Update("kata_sandi", string(hashedPassword)).Error
}

// ALAMAT

// Menambah alamat baru
func CreateAlamat(input models.Alamat) (models.Alamat, error) {
	err := config.DB.Create(&input).Error
	return input, err
}

// Mengambil semua alamat milik user yang sedang login
func GetMyAlamat(userID uint) ([]models.Alamat, error) {
	var alamat []models.Alamat
	err := config.DB.Where("id_user = ?", userID).Find(&alamat).Error
	return alamat, err
}

// Memperbarui alamat
func UpdateAlamat(id uint, userID uint, input models.Alamat) error {
	return config.DB.Model(&models.Alamat{}).Where("id = ? AND id_user = ?", id, userID).Updates(input).Error
}

// Menghapus alamat
func DeleteAlamat(id uint, userID uint) error {
	return config.DB.Where("id = ? AND id_user = ?", id, userID).Delete(&models.Alamat{}).Error
}
