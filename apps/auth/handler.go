package auth

import (
	"encoding/json"
	"evermos/models"
	"net/http"
)

// REGISTER
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var input models.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := Register(input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  false,
			"message": "Gagal Register Pengguna",
			"errors":  []string{err.Error()},
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "Berhasil membuat data",
		"data":    user,
	})
}

// LOGIN
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		NoTelp    string `json:"notelp"`
		KataSandi string `json:"kata_sandi"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result, err := Login(input.NoTelp, input.KataSandi)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  false,
			"message": "Gagal Login",
			"errors":  []string{err.Error()},
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "Berhasil Login",
		"data":    result,
	})
}

// LOGOUT
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Untuk stateless JWT, logout cukup memberikan respons sukses
	// dan membiarkan client menghapus tokennya sendiri
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "Berhasil Logout",
		"data":    nil,
	})
}
