package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ArteShow/Donut-SMP-Price-Searcher/internal/core"
)

func GetAveregPrice(w http.ResponseWriter, r *http.Request) {
	var req GetAveregPriceRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if err = json.Unmarshal(body, &req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	price := core.CalculateAveregPrice(req.Token)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(GetAveregPriceResponse{Price: price}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}