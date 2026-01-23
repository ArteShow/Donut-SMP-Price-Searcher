package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ArteShow/Donut-SMP-Price-Searcher/internal/core"
)

func GetAveregPrice(w http.ResponseWriter, r *http.Request) {
	log.Println("[HANDLER] /price request received")

	var req GetAveregPriceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("[HANDLER] invalid JSON:", err)
		http.Error(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	log.Printf("[HANDLER] token=%s item=%s\n", req.Token, req.Item)

	price, err := core.CalculateAveregPrice(req.Token, req.Item)
	if price == 0 {
		http.Error(w, "item not found", http.StatusNotFound)
		return
	}

	if err != nil {
		log.Println("[HANDLER] calculate error:", err)
		http.Error(w, "failed to get price: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("[HANDLER] calculated price=%d\n", price)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(GetAveregPriceResponse{Price: price})
}
