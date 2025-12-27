# Mini Project - Golang Backend Evermos

Repositori ini adalah program backend menggunakan Golang untuk aplikasi Evermos. Evermos merupakan platform social commerce reseller untuk menjual produk-produk Muslim Indonesia. Pada program ini terdapat beberapa fitur, diantaranya:
- Register dan Login untuk User.
- Pengelolaan akun dan toko.
- Pengelolaan produk dan alamat.
- Pengelolaan kategori bagi admin.
- Transaksi dengan log riwayat.

---

## Software yang Digunakan
- **Golang** sebagai bahasa pemrograman.
- **Framework Fiber** sebagai framework RESTAPI.
- **MySQL** sebagai database.
- **Testing** (seperti Postman, HTTPie, Command Line, dan sebagainya).

---

## Cara Menjalankan Program
### 1. Clone repositori
```
https://github.com/muhzarfan/miniproject-evermos.git
cd miniproject-evermos
```

### 2. Instal Framework dan Dependensi
```
go get github.com/gofiber/fiber/v2
```
```
go get gorm.io/gorm
```
```
go get gorm.io/driver/mysql
```
```
go get github.com/golang-jwt/jwt/v4
```
```
go get golang.org/x/crypto/bcrypt
```

### 3. Pilih Database yang Digunakan
Program ini menggunakan database MySQL di lokal dan sesuaikan koneksi database pada file `config/database.go`.

### 3. Jalankan Program
Sebelum jalankan program, pastikan membuat folder `uploads` untuk menampung file gambar dan jalankan:
```
go run main.go
```

---

## Dokumentasi API
**Note:**
Beberapa endpoint membutuhkan `token` untuk melakukan request yang didapat dari body response setelah login pengguna.
### 1. Register User
Request:
```http
POST http://localhost:8000/api/auth/login
```
Body:
```json
{
    "nama": "John Doe",
    "kata_sandi": "inipasword123",
    "notelp": "08123467890",
    "tanggal_lahir": "2003-05-26",
    "jenis_kelamin": "Laki-laki",
    "tentang": "Saya adalah pengusaha pakaian muslim.",
    "pekerjaan": "Wirausaha",
    "email": "johndoe@gmail.com",
    "id_provinsi": "32",
    "id_kota": "3276"
}
```
Untuk `id_provinsi` dan `id_kota` dapat diisi sesuai dari [API Statis Wilayah Indonesia](https://www.emsifa.com/api-wilayah-indonesia/). Toko milik user akan otomatis dibuat saat register.

### 2. Login User
Request:
```http
POST http://localhost:8000/api/auth/login
```
Body:
```json
{
    "notelp": "08123467890",
    "kata_sandi": "inipassword123"
}
```
Token akan muncul di response yang expired setelah 7 hari.

### 3. Service Akun User
**a. Get Profile** <br>
Request:
```http
GET http://localhost:8000/api/user/me
```
Header:
| Key | Value |
|-----|-------|
| token | token-acak-login |

**b. Edit Profile** <br>
Request:
```http
PUT http://localhost:8000/api/user/me
```
Body (tidak harus isi semua yang diubah):
```json
{
    "nama": "John Steve Doe",
    "tanggal_lahir": "2003-08-17",
    "pekerjaan": "Wirausaha",
    "id_provinsi": "31",
    "id_kota": "3171"
}
```

**c. Ubah Password** <br>
Request:
```http
PUT http://localhost:8000/api/user/change-password
```
Body:
```json
{
    "password_baru": "passwordBaru123"
}
```
Header:
| Key | Value |
|-----|-------|
| token | token-acak-login |

### 4. Service Toko
**a. Get Toko** <br>
Request:
```http
GET http://localhost:8000/api/toko/my
```
Header:
| Key | Value |
|-----|-------|
| token | token-acak-login |

**b. Edit Toko** <br>
Request:
```http
PUT http://localhost:8000/api/toko/my
```
Request (form data):
| Key | Tipe | Value |
|-----|------|-----------|
| nama_toko | Text | Nama toko misal: John Doe Official Store |
| url_toko | File | File gambar untuk profil toko |

Header:
| Key | Value |
|-----|-------|
| token | token-acak-login ||

### 5. Service Alamat
**a. Get Alamat** <br>
Request:
```http
GET http://localhost:8000/api/user/alamat
```
Header:
| Key | Value |
|-----|-------|
| token | token-acak-login |

**b. Create Alamat** <br>
Request:
```http
POST http://localhost:8000/api/user/alamat
```
Body:
```json
{
    "judul_alamat": "Gudang Reseller",
    "nama_penerima": "Rusdi",
    "no_telp": "086667778812",
    "detail_alamat": "Jl. Merapi No.3A, Cilandak, Jakarta Selatan"
}
```
Header:
| Key | Value |
|-----|-------|
| token | token-acak-login |

**c. Edit Alamat** <br>
Request (`id` disesuaikan dengan alamat yang ingin diubah):
```http
PUT http://localhost:8000/api/user/alamat/{:id}
```
Body:
```json
{
    "judul_alamat": "Head Office",
    "nama_penerima": "Gwen",
    "no_telp": "086667778812",
    "detail_alamat": "Jl. Bromo No.15, Pasar Minggu, Jakarta Selatan"
}
```
Header:
| Key | Value |
|-----|-------|
| token | token-acak-login |

**d. Delete Alamat** <br>
Request (`id` disesuaikan dengan alamat yang ingin dihapus):
```http
DELETE http://localhost:8000/api/user/alamat/{:id}
```

### 6. Service Kategori
**a. Get Kategori** <br>
Request:
```http
GET http://localhost:8000/api/category
```
Header:
| Key | Value |
|-----|-------|
| token | token-acak-login |

**b. Create Kategori (Hanya Admin)** <br>
Request:
```http
POST http://localhost:8000/api/category
```
Body:
```json
{
    "nama_category": "Baju"
}
```
Header:
| Key | Value |
|-----|-------|
| token | token-acak-login |

**c. Edit Kategori (Hanya Admin)** <br>
Request (`id` disesuaikan dengan kategori yang ingin diubah):
```http
PUT http://localhost:8000/api/category/{:id}
```
Body:
```json
{
    "nama_category": "Celana"
}
```
Header:
| Key | Value |
|-----|-------|
| token | token-acak-login |

**d. Delete Kategori (Hanya Admin)** <br>
Request (`id` disesuaikan dengan kategori yang ingin dihapus):
```http
DELETE http://localhost:8000/api/category/{:id}
```

### 7. Service Produk
**a. Get Kategori** <br>
Request:
```http
GET http://localhost:8000/api/produk
```
Header:
| Key | Value |
|-----|-------|
| token | token-acak-login |

**b. Create Produk** <br>
Request:
```http
POST http://localhost:8000/api/produk
```
Body (`id_category` disesuaikan dengan kategori yang tersedia):
```json
{
    "nama_produk": "Baju Koko Pria",
    "slug": "baju-koko-pria",
    "harga_reseller": "180000",
    "harga_konsumen": "200000",
    "stok": 50,
    "deskripsi": "Baju koko pria kasual yang cocok untuk acara dan ibadah.",
    "id_category": 1
}
```
Header:
| Key | Value |
|-----|-------|
| token | token-acak-login |

**c. Edit Produk** <br>
Request (`id` sesuaikan dengan produk yang ingin diubah):
```http
PUT http://localhost:8000/api/produk/{:id}
```
Body:
```json
{
    "nama_produk": "Baju Gamis Pria",
    "slug": "baju-gamis-pria",
    "harga_reseller": "180000",
    "harga_konsumen": "200000",
    "stok": 100,
    "deskripsi": "Baju gamis pria kasual yang cocok untuk acara dan ibadah.",
    "id_category": 2
}
```
Header:
| Key | Value |
|-----|-------|
| token | token-acak-login |

**d. Delete Produk** <br>
Request (`id` disesuaikan dengan produk yang ingin dihapus):
```http
DELETE http://localhost:8000/api/produk/{:id}
```

### 7. Service Transaksi
**a. Get Transaksi** <br>
Request:
```http
GET http://localhost:8000/api/trx
```
Header:
| Key | Value |
|-----|-------|
| token | token-acak-login |

**b. Create Transaksi** <br>
Request:
```http
POST http://localhost:8000/api/trx
```
Body:
```json
{
    "alamat_pengiriman": 1,
    "method_bayar": "Transffer Bank",
    "details": [
        {
            "id_log_produk": 2,
            "kuantitas": 20
        }
    ]
}
```
- `alamat_pengiriman` disesuaikan dengan id alamat tujuan.
- `id_log_produk` disesuaikan dengan id produk.
- Program akan otomatis mengurangi stok produk setelah melakukan transaksi.
- Hasil invoice akan disimpan pada tabel `detail_trx` dan produk akan terisi pada tabel `log_produk` ketika melakukan transaksi. 
