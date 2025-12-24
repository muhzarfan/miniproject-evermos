package wilayah

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Fungsi pembantu untuk melakukan fetch ke API
func fetchEmsifa(url string) ([]map[string]interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	return data, err
}

func GetProvinces() ([]map[string]interface{}, error) {
	return fetchEmsifa("https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json")
}

func GetRegencies(provID string) ([]map[string]interface{}, error) {
	url := fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/regencies/%s.json", provID)
	return fetchEmsifa(url)
}
