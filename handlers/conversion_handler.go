package handlers

import (
	"encoding/json"
	"myrestapi/models"
	"myrestapi/utils"
	"net/http"
)

func HandleConversion(w http.ResponseWriter, r *http.Request) {
	var req models.ConversionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pounds := utils.KgToPounds(req.Kilograms)
	res := models.ConversionResponse{Pounds: pounds}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
