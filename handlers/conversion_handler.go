// handlers/conversion_handler.go

package handlers

import (
	"encoding/json"
	"myrestapi/models"
	"myrestapi/utils"
	"net/http"
)

// HandleConversion handles the conversion request.
func HandleConversion(w http.ResponseWriter, r *http.Request) {
	var req models.ConversionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pounds := utils.KgToPounds(req.Kilograms) // Use KgToPounds function from utils package
	res := models.ConversionResponse{Pounds: pounds}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
