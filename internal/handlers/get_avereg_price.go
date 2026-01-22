package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ArteShow/Donut-SMP-Price-Searcher/internal/core"
)

func GetAveregPrice(w http.ResponseWriter, r *http.Request) {
	var req GetAveregPriceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	price, err := core.CalculateAveregPrice(req.Token, req.Item)
	if err != nil {
		http.Error(w, "failed to get price: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(GetAveregPriceResponse{Price: price})
}

