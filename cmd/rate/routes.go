package rate

import (
	"encoding/json"
	"net/http"
)

var rateService = &RateService{}

func RateHandler(w http.ResponseWriter, r *http.Request) {
	rateName := r.URL.Query().Get("name")

	data, err := rateService.GetRateData(rateName)
	if err != nil {
		http.Error(w, "Failed to fetch rate data", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to format rate data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonData)
}
