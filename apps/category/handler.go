package category

import (
	"encoding/json"
	"evermos/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Membuat kategori baru (POST)
func CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var input models.Category
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	res, err := CreateCategory(input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  false,
			"message": "Gagal membuat kategori",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "Berhasil membuat data",
		"data":    res.ID,
	})
}

// Mengambil semua kategori (GET)
func GetAllCategoryHandler(w http.ResponseWriter, r *http.Request) {
	res, err := GetAllCategories()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"status": false, "message": "Gagal mengambil data"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "Berhasil mendapatkan data",
		"data":    res,
	})
}

// Memperbarui kategori (PUT)
func UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var input models.Category
	json.NewDecoder(r.Body).Decode(&input)

	err := UpdateCategory(uint(id), input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"status": false, "message": "Gagal update kategori"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "Berhasil memperbarui data",
		"data":    "",
	})
}

// Menghapus kategori (DELETE)
func DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	err := DeleteCategory(uint(id))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"status": false, "message": "Gagal menghapus kategori"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "Berhasil menghapus data",
		"data":    "",
	})
}
