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
	case "KgToPounds":
		result = fmt.Sprintf("%.2f", utils.KgToPounds(req.Value)) + "lbs"
	case "PoundsToKg":
		result = fmt.Sprintf("%.2f", utils.PoundsToKg(req.Value)) + "kg"
	case "KmToMiles":
		result = fmt.Sprintf("%.2f", utils.KmToMiles(req.Value)) + "miles"
	case "MilesToKm":
		result = fmt.Sprintf("%.2f", utils.MilesToKm(req.Value)) + "km"
	case "CmToFeetInches":
		feet, inches := utils.CmToFeetInches(req.Value)
		result = fmt.Sprintf("%d feet and %.2f inches", feet, inches)
	default:
		http.Error(w, "Invalid conversion type", http.StatusBadRequest)
		return
	}

	res := models.ConversionResponse{Value: result}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
