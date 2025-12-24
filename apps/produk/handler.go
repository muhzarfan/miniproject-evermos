package produk

import (
	"encoding/json"
	"evermos/config"
	"evermos/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Membuat produk baru (POST)
func CreateProdukHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(float64)

	// Cari ID Toko berdasarkan User ID
	var toko models.Toko
	config.DB.Where("id_user = ?", userID).First(&toko)

	var input models.Produk
	json.NewDecoder(r.Body).Decode(&input)

	input.IDToko = toko.ID

	res, err := CreateProduk(input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"status": false, "message": "Gagal buat produk"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"status": true, "message": "Berhasil membuat data", "data": res.ID})
}

// Mengambil data produk (GET)
func GetAllProdukHandler(w http.ResponseWriter, r *http.Request) {
	// Mengambil query parameter dari URL
	nama := r.URL.Query().Get("nama")
	categoryID := r.URL.Query().Get("id_category")

	res, err := GetAllProduk(nama, categoryID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  false,
			"message": "Gagal mengambil data",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "Berhasil mengambil semua data",
		"data":    res,
	})
}

// Mengambil produk berdasarkan ID (GET)
func GetProdukByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	res, err := GetProdukByID(uint(id))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{"status": false, "message": "Produk tidak ditemukan"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"status": true, "message": "Berhasil mengambil data berdasarkan ID", "data": res})
}

// Memperbarui produk (PUT)
func UpdateProdukHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	userID := r.Context().Value("user_id").(float64)

	// Cari ID Toko pemilik
	var toko models.Toko
	config.DB.Where("id_user = ?", userID).First(&toko)

	var input models.Produk
	json.NewDecoder(r.Body).Decode(&input)

	err := UpdateProduk(uint(id), toko.ID, input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "Berhasil memperbarui data",
		"data":    "",
	})
}

// Menghapus data produk (DELETE)
func DeleteProdukHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	userID := r.Context().Value("user_id").(float64)

	var toko models.Toko
	config.DB.Where("id_user = ?", userID).First(&toko)

	err := DeleteProduk(uint(id), toko.ID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"status": false, "message": "Gagal hapus produk"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"status": true, "message": "Berhasil menghapus data", "data": ""})
}
