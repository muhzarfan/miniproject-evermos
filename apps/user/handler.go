package user

import (
	"encoding/json"
	"evermos/apps/wilayah"
	"evermos/models"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Mengambil Data User yang Login
func GetMyProfileHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(float64)
	userData, err := GetMyProfile(uint(userID))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  false,
			"message": "User tidak ditemukan",
		})
		return
	}

	response := UserProfileResponse{User: userData}

	// Ambil Data API Wilayah untuk Provinsi dan ota
	if userData.IDProvinsi != "" {
		provList, _ := wilayah.GetProvinces()
		for _, p := range provList {
			if fmt.Sprintf("%v", p["id"]) == userData.IDProvinsi {
				response.NamaProvinsi = p["name"].(string)
				break
			}
		}
	}

	if userData.IDKota != "" {
		kotaList, _ := wilayah.GetRegencies(userData.IDProvinsi)
		for _, k := range kotaList {
			if fmt.Sprintf("%v", k["id"]) == userData.IDKota {
				response.NamaKota = k["name"].(string)
				break
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "Berhasil mengambil data",
		"data":    response,
	})
}

type UserProfileResponse struct {
	models.User
	NamaProvinsi string `json:"nama_provinsi"`
	NamaKota     string `json:"nama_kota"`
}

// Perbarui Akun User yang Login (PUT)
func UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(float64)

	var input models.User
	json.NewDecoder(r.Body).Decode(&input)

	err := UpdateProfile(uint(userID), input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"status": false, "message": "Gagal update profil"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "Update profil berhasil",
	})
}

// Ubah Password User (PUT)
func ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(float64)

	var input struct {
		PasswordBaru string `json:"password_baru"`
	}
	json.NewDecoder(r.Body).Decode(&input)

	if len(input.PasswordBaru) < 6 {
		http.Error(w, "Password minimal 6 karakter", http.StatusBadRequest)
		return
	}

	err := ChangePassword(uint(userID), input.PasswordBaru)
	if err != nil {
		http.Error(w, "Gagal mengganti password", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"status": true, "message": "Password berhasil diganti"})
}

// Membuat alamat baru (POST)
func CreateAlamatHandler(w http.ResponseWriter, r *http.Request) {
	var input models.Alamat
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Mengambil userID dari context (dari AuthMiddleware)
	userIDFloat := r.Context().Value("user_id").(float64)
	input.IDUser = uint(userIDFloat)

	res, err := CreateAlamat(input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  false,
			"message": "Gagal menambah alamat",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "Berhasil menyimpan data",
		"data":    res.ID,
	})
}

// Mengambil data alamat (GET)
func GetMyAlamatHandler(w http.ResponseWriter, r *http.Request) {
	userIDFloat := r.Context().Value("user_id").(float64)

	res, err := GetMyAlamat(uint(userIDFloat))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  false,
			"message": "Gagal mengambil data alamat",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "Berhasil mendapat data",
		"data":    res,
	})
}

// Memperbarui alamat (PUT)
func UpdateAlamatHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var input models.Alamat
	json.NewDecoder(r.Body).Decode(&input)

	userID := r.Context().Value("user_id").(float64)

	var addrID uint
	fmt.Sscanf(id, "%d", &addrID)

	err := UpdateAlamat(addrID, uint(userID), input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"status": false, "message": "Gagal update atau data tidak ditemukan"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"status": true, "message": "Berhasil update data.", "data": ""})
}

// Menghapus alamat (DELETE)
func DeleteAlamatHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	userID := r.Context().Value("user_id").(float64)

	var addrID uint
	fmt.Sscanf(id, "%d", &addrID)

	err := DeleteAlamat(addrID, uint(userID))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"status": false, "message": "Gagal menghapus data"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"status": true, "message": "Berhasil menghapus data.", "data": ""})
}
