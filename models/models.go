package models

import (
	"time"
)

// Model User
type User struct {
	ID           uint      `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Nama         string    `gorm:"type:varchar(255);column:nama" json:"nama"`
	KataSandi    string    `gorm:"type:varchar(255);column:kata_sandi" json:"kata_sandi"`
	Notelp       string    `gorm:"type:varchar(255);unique;column:notelp" json:"notelp"`
	TanggalLahir string    `gorm:"type:date;column:tanggal_lahir" json:"tanggal_lahir"`
	JenisKelamin string    `gorm:"type:varchar(255);column:jenis_kelamin" json:"jenis_kelamin"`
	Tentang      string    `gorm:"type:text;column:tentang" json:"tentang"`
	Pekerjaan    string    `gorm:"type:varchar(255);column:pekerjaan" json:"pekerjaan"`
	Email        string    `gorm:"type:varchar(255);column:email" json:"email"`
	IDProvinsi   string    `gorm:"type:varchar(255);column:id_provinsi" json:"id_provinsi"`
	IDKota       string    `gorm:"type:varchar(255);column:id_kota" json:"id_kota"`
	IsAdmin      bool      `gorm:"type:boolean;default:false;column:isAdmin" json:"is_admin"`
	UpdatedAt    time.Time `gorm:"type:date;column:updated_at" json:"updated_at"`
	CreatedAt    time.Time `gorm:"type:date;column:created_at" json:"created_at"`
	Alamat       []Alamat  `gorm:"foreignKey:IDUser;column:id"`
}

func (User) TableName() string {
	return "User"
}

// Model Toko
type Toko struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	IDUser    uint      `gorm:"column:id_user" json:"id_user"`
	NamaToko  string    `gorm:"type:varchar(255);column:nama_toko" json:"nama_toko"`
	UrlFoto   string    `gorm:"type:varchar(255);column:url_foto" json:"url_foto"`
	UpdatedAt time.Time `gorm:"type:date;column:updated_at" json:"updated_at"`
	CreatedAt time.Time `gorm:"type:date;column:created_at" json:"created_at"`
}

func (Toko) TableName() string {
	return "toko"
}

// Model Alamat
type Alamat struct {
	ID           uint      `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	IDUser       uint      `gorm:"column:id_user" json:"id_user"`
	JudulAlamat  string    `gorm:"type:varchar(255);column:judul_alamat" json:"judul_alamat"`
	NamaPenerima string    `gorm:"type:varchar(255);column:nama_penerima" json:"nama_penerima"`
	NoTelp       string    `gorm:"type:varchar(255);column:no_telp" json:"no_telp"`
	DetailAlamat string    `gorm:"type:varchar(255);column:detail_alamat" json:"detail_alamat"`
	UpdatedAt    time.Time `gorm:"type:date;column:updated_at" json:"updated_at"`
	CreatedAt    time.Time `gorm:"type:date;column:created_at" json:"created_at"`
}

func (Alamat) TableName() string {
	return "alamat"
}

// Model Kategori
type Category struct {
	ID           uint      `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	NamaCategory string    `gorm:"type:varchar(255);column:nama_category" json:"nama_category"`
	UpdatedAt    time.Time `gorm:"type:date;column:updated_at" json:"updated_at"`
	CreatedAt    time.Time `gorm:"type:date;column:created_at" json:"created_at"`
}

func (Category) TableName() string {
	return "category"
}

// Model Produk
type Produk struct {
	ID            uint      `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	NamaProduk    string    `gorm:"type:varchar(255);column:nama_produk" json:"nama_produk"`
	Slug          string    `gorm:"type:varchar(255);column:slug" json:"slug"`
	HargaReseller string    `gorm:"type:varchar(255);column:harga_reseller" json:"harga_reseller"`
	HargaKonsumen string    `gorm:"type:varchar(255);column:harga_konsumen" json:"harga_konsumen"`
	Stok          int       `gorm:"type:int;column:stok" json:"stok"`
	Deskripsi     string    `gorm:"type:text;column:deskripsi" json:"deskripsi"`
	IDToko        uint      `gorm:"column:id_toko" json:"id_toko"`
	IDCategory    uint      `gorm:"column:id_category" json:"id_category"`
	UpdatedAt     time.Time `gorm:"type:date;column:updated_at" json:"updated_at"`
	CreatedAt     time.Time `gorm:"type:date;column:created_at" json:"created_at"`
}

func (Produk) TableName() string {
	return "produk"
}

// Model Transaksi dan Detail
type Trx struct {
	ID               uint        `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	IDUser           uint        `gorm:"column:id_user" json:"id_user"`
	AlamatPengiriman uint        `gorm:"column:alamat_pengiriman" json:"alamat_pengiriman"`
	HargaTotal       int         `gorm:"column:harga_total" json:"harga_total"`
	KodeInvoice      string      `gorm:"type:varchar(255);column:kode_invoice" json:"kode_invoice"`
	MethodBayar      string      `gorm:"type:varchar(255);column:method_bayar" json:"method_bayar"`
	UpdatedAt        time.Time   `gorm:"type:date;column:updated_at" json:"updated_at"`
	CreatedAt        time.Time   `gorm:"type:date;column:created_at" json:"created_at"`
	Details          []DetailTrx `gorm:"foreignKey:IDTrx" json:"details"`
}

func (Trx) TableName() string {
	return "trx"
}

type DetailTrx struct {
	ID          uint      `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	IDTrx       uint      `gorm:"column:id_trx" json:"id_trx"`
	IDLogProduk uint      `gorm:"column:id_log_produk" json:"id_log_produk"`
	IDToko      uint      `gorm:"column:id_toko" json:"id_toko"`
	Kuantitas   int       `gorm:"column:kuantitas" json:"kuantitas"`
	HargaTotal  int       `gorm:"column:harga_total" json:"harga_total"`
	UpdatedAt   time.Time `gorm:"type:date;column:updated_at" json:"updated_at"`
	CreatedAt   time.Time `gorm:"type:date;column:created_at" json:"created_at"`
}

func (DetailTrx) TableName() string {
	return "detail_trx"
}

// Model Riwayat Transaksi
type LogProduk struct {
	ID            uint      `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	IDProduk      uint      `gorm:"column:id_produk" json:"id_produk"`
	NamaProduk    string    `gorm:"type:varchar(255);column:nama_produk" json:"nama_produk"`
	Slug          string    `gorm:"type:varchar(255);column:slug" json:"slug"`
	HargaReseller string    `gorm:"type:varchar(255);column:harga_reseller" json:"harga_reseller"`
	HargaKonsumen string    `gorm:"type:varchar(255);column:harga_konsumen" json:"harga_konsumen"`
	Deskripsi     string    `gorm:"type:text;column:deskripsi" json:"deskripsi"`
	IDToko        uint      `gorm:"column:id_toko" json:"id_toko"`
	IDCategory    uint      `gorm:"column:id_category" json:"id_category"`
	UpdatedAt     time.Time `gorm:"type:date;column:updated_at" json:"updated_at"`
	CreatedAt     time.Time `gorm:"type:date;column:created_at" json:"created_at"`
}

func (LogProduk) TableName() string {
	return "log_produk"
}
