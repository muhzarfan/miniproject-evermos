package toko

import (
	"encoding/json"
	"evermos/models"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// Mengambil data toko (GET)
func GetMyTokoHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(float64)

	res, err := GetMyToko(uint(userID))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{"status": false, "message": "Toko tidak ditemukan"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "Berhasil mendapat data.",
		"data":    res,
	})
}

// Memperbarui toko (PUT)
func UpdateTokoHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tokoIDStr := params["id_toko"]
	userID := r.Context().Value("user_id").(float64)

	// Konversi ID Toko ke uint
	id, err := strconv.ParseUint(tokoIDStr, 10, 32)
	if err != nil {
		http.Error(w, "ID Toko tidak valid", http.StatusBadRequest)
		return
	}

	// Multipart Form (File max 10MB)
	err = r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "File terlalu besar", http.StatusBadRequest)
		return
	}

	// Ambil file foto
	file, handler, err := r.FormFile("url_foto")
	var fileName string
	if err == nil {
		defer file.Close()
		// Pastikan folder 'uploads' ada
		os.MkdirAll("uploads", os.ModePerm)
		fileName = "uploads/" + handler.Filename

		f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
		if err == nil {
			defer f.Close()
			io.Copy(f, file)
		}
	}

	namaToko := r.FormValue("nama_toko")

	input := models.Toko{
		NamaToko: namaToko,
		UrlFoto:  fileName,
	}

	// Panggil service
	err = UpdateToko(uint(id), uint(userID), input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		// Jika error mengandung kata 'akses', kirim 403 Forbidden
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
		"data":    "Perbarui toko berhasil",
	})
}
