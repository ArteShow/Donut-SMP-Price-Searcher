package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ArteShow/Donut-SMP-Price-Searcher/internal/service"
)

func GetLowestPricePerItem(w http.ResponseWriter, r *http.Request) {
	var req GetLowestPricePerItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	price, err := service.GetLowestAveragePricePerItem(req.ID, req.Token, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if price == 0.0 {
		http.Error(w, "item not found", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(GetLowestPricePerItemResponse{Price: price})
}