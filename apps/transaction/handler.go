package transaction

import (
	"encoding/json"
	"evermos/models"
	"net/http"
)

// Membuat Transaksi (POST)
func CreateTrxHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(float64)

	var input models.Trx
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	res, err := CreateTransaction(uint(userID), input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"status": false, "message": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "Berhasil membuat invoice",
		"data":    res,
	})
}

// Mengambil Data Transaksi (GET)
func GetAllTrxHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(float64)

	namaProduk := r.URL.Query().Get("nama_produk")
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	res, err := GetAllTransaction(uint(userID), namaProduk, startDate, endDate)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"status": false, "message": "Gagal mengambil riwayat"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "Berhasil mengambil data",
		"data":    res,
	})
}
