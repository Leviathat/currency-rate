package rate

import (
	"encoding/json"
	"net/http"
)

var rateService = &RateService{}

func RateHandler(w http.ResponseWriter, r *http.Request) {
	// Extract rate name from request parameters, you may want to add validation here
	rateName := r.URL.Query().Get("name")

	// Call the RateService to get data
	data, err := rateService.GetRateData(rateName)
	if err != nil {
		http.Error(w, "Failed to fetch rate data", http.StatusInternalServerError)
		return
	}

	// Marshal data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to format rate data", http.StatusInternalServerError)
		return
	}

	// Set Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(jsonData)
}
