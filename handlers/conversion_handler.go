package handlers

import (
	"encoding/json"
	"fmt"
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

	var result string
	switch req.Type {
	case "kgToPounds":
		result = fmt.Sprintf("%.2f", utils.KgToPounds(req.Value)) + "lbs"
	case "poundsToKg":
		result = fmt.Sprintf("%.2f", utils.PoundsToKg(req.Value)) + "kg"
	default:
		http.Error(w, "Invalid conversion type", http.StatusBadRequest)
		return
	}

	res := models.ConversionResponse{Value: result}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
