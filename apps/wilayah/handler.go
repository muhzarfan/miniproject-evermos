package wilayah

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetProvincesHandler(w http.ResponseWriter, r *http.Request) {
	res, _ := GetProvinces()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func GetRegenciesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	res, _ := GetRegencies(params["id_provinsi"])
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
