package toko

import (
	"errors"
	"evermos/config"
	"evermos/models"
)

// Ambil data toko milik user yang login
func GetMyToko(userID uint) (models.Toko, error) {
	var toko models.Toko
	err := config.DB.Where("id_user = ?", userID).First(&toko).Error
	return toko, err
}

// Update profil toko
func UpdateToko(tokoID uint, userID uint, input models.Toko) error {
	updateData := make(map[string]interface{})
	if input.NamaToko != "" {
		updateData["nama_toko"] = input.NamaToko
	}
	if input.UrlFoto != "" {
		updateData["url_foto"] = input.UrlFoto
	}

	result := config.DB.Model(&models.Toko{}).
		Where("id = ? AND id_user = ?", tokoID, userID).
		Updates(updateData)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("invalid: Toko tidak ditemukan atau Anda tidak memiliki akses untuk mengubah toko ini")
	}

	return nil
}
