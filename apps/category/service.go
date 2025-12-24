package category

import (
	"evermos/config"
	"evermos/models"
)

// Menambah kategori baru (Hanya Admin)
func CreateCategory(input models.Category) (models.Category, error) {
	err := config.DB.Create(&input).Error
	return input, err
}

// Mengambil semua daftar kategori
func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	err := config.DB.Find(&categories).Error
	return categories, err
}

// Update kategori berdasarkan ID
func UpdateCategory(id uint, input models.Category) error {
	return config.DB.Model(&models.Category{}).Where("id = ?", id).Updates(input).Error
}

// Hapus kategori berdasarkan ID
func DeleteCategory(id uint) error {
	return config.DB.Where("id = ?", id).Delete(&models.Category{}).Error
}
