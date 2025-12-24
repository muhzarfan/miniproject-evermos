package transaction

import (
	"errors"
	"evermos/config"
	"evermos/models"
	"fmt"
	"strconv"
	"time"
)

// Membuat Transaksi
func CreateTransaction(userID uint, input models.Trx) (models.Trx, error) {
	tx := config.DB.Begin()

	var totalHargaInvoice int
	for i, detail := range input.Details {
		var produk models.Produk
		// Dapat data produk
		if err := tx.First(&produk, detail.IDLogProduk).Error; err != nil {
			tx.Rollback()
			return input, errors.New("message: Produk tidak ditemukan")
		}

		// Cek stok produk
		if produk.Stok < detail.Kuantitas {
			tx.Rollback()
			return input, fmt.Errorf("message: Stok produk %s tidak mencukupi", produk.NamaProduk)
		}

		// Salin data ke log produk
		logProduk := models.LogProduk{
			IDProduk:      produk.ID,
			NamaProduk:    produk.NamaProduk,
			Slug:          produk.Slug,
			HargaReseller: produk.HargaReseller,
			HargaKonsumen: produk.HargaKonsumen,
			Deskripsi:     produk.Deskripsi,
			IDToko:        produk.IDToko,
			IDCategory:    produk.IDCategory,
		}
		if err := tx.Create(&logProduk).Error; err != nil {
			tx.Rollback()
			return input, errors.New("gagal mencatat log produk")
		}

		// Hitung harga produk sesuai database
		hargaSatuan, _ := strconv.Atoi(produk.HargaKonsumen)
		itemTotalHarga := hargaSatuan * detail.Kuantitas

		// Kurangi stok
		tx.Model(&produk).Update("stok", produk.Stok-detail.Kuantitas)

		// Update data detail
		input.Details[i].IDLogProduk = logProduk.ID
		input.Details[i].IDToko = produk.IDToko
		input.Details[i].HargaTotal = itemTotalHarga

		totalHargaInvoice += itemTotalHarga
	}

	input.IDUser = userID
	input.HargaTotal = totalHargaInvoice
	input.KodeInvoice = fmt.Sprintf("INV/%d/%s", userID, time.Now().Format("20060102150405"))

	if err := tx.Create(&input).Error; err != nil {
		tx.Rollback()
		return input, err
	}

	tx.Commit()
	return input, nil
}

// Mengambil Data Transaksi
func GetAllTransaction(userID uint, namaProduk string, startDate string, endDate string) ([]models.Trx, error) {
	var transactions []models.Trx

	// Ambil data transaksi milik user yang login
	query := config.DB.Preload("Details").Where("id_user = ?", userID)

	// Filter berdasarkan Rentang Tanggal
	if startDate != "" && endDate != "" {
		query = query.Where("created_at BETWEEN ? AND ?", startDate, endDate)
	}

	// Filter berdasarkan Nama Produk
	if namaProduk != "" {
		query = query.Joins("JOIN detail_trx ON detail_trx.id_trx = trx.id").
			Joins("JOIN log_produk ON log_produk.id = detail_trx.id_log_produk").
			Where("log_produk.nama_produk LIKE ?", "%"+namaProduk+"%").
			Group("trx.id")
	}

	err := query.Find(&transactions).Error
	return transactions, err
}
