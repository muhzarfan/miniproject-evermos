package produk

import (
	"errors"
	"evermos/config"
	"evermos/models"
)

// Membuat data produk baru
func CreateProduk(input models.Produk) (models.Produk, error) {
	err := config.DB.Create(&input).Error
	return input, err
}

// Mengambil data produk
func GetAllProduk(nama string, catID string) ([]models.Produk, error) {
	var produk []models.Produk
	query := config.DB
	if nama != "" {
		query = query.Where("nama_produk LIKE ?", "%"+nama+"%")
	}
	if catID != "" {
		query = query.Where("id_category = ?", catID)
	}
	err := query.Find(&produk).Error
	return produk, err
}

// Ambil detail satu produk
func GetProdukByID(id uint) (models.Produk, error) {
	var produk models.Produk
	err := config.DB.First(&produk, id).Error
	return produk, err
}

// Update produk milik toko sendiri
func UpdateProduk(id uint, tokoID uint, input models.Produk) error {
	result := config.DB.Model(&models.Produk{}).
		Where("id = ? AND id_toko = ?", id, tokoID).
		Updates(input)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("invalid: Produk tidak ditemukan atau Anda bukan pemilik toko ini")
	}

	return nil
}

// Hapus produk milik toko sendiri
func DeleteProduk(id uint, tokoID uint) error {
	return config.DB.Where("id = ? AND id_toko = ?", id, tokoID).Delete(&models.Produk{}).Error
}
