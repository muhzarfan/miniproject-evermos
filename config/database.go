package config

import (
	"evermos/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Koneksi ke Database
	dsn := "root:@tcp(127.0.0.1:3306)/evermos_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal terhubung ke database")
	}

	// Buat Tabel Otomatis
	db.AutoMigrate(
		&models.User{},
		&models.Toko{},
		&models.Alamat{},
		&models.Category{},
		&models.Produk{},
		&models.Trx{},
		&models.DetailTrx{},
		&models.LogProduk{},
	)

	DB = db
	fmt.Println("Database terhubung dan migrasi berhasil!")
}
