package core

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/ArteShow/Donut-SMP-Price-Searcher/internal/models"
)

const BaseURL = "https://api.donutsmp.net"

func SendRequest(token, endpoint string) (models.ListAuctionPageResponse, error) {
	log.Println("[HTTP] GET", BaseURL+endpoint)

	Request, err := json.Marshal(models.ListAuctionPageRequest{
		Sort: "lowest_price",
	})
	if err != nil {
		log.Println("[HTTP] marshal error:", err)
		return models.ListAuctionPageResponse{}, err
	}

	req, err := http.NewRequest(http.MethodGet, BaseURL+endpoint, bytes.NewReader(Request))
	if err != nil {
		log.Println("[HTTP] request build error:", err)
		return models.ListAuctionPageResponse{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("[HTTP] request failed:", err)
		return models.ListAuctionPageResponse{}, err
	}
	defer resp.Body.Close()

	log.Println("[HTTP] status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("[HTTP] read body error:", err)
		return models.ListAuctionPageResponse{}, err
	}

	log.Println("[HTTP] body length:", len(body))

	var res models.ListAuctionPageResponse
	if err = json.Unmarshal(body, &res); err != nil {
		log.Println("[HTTP] unmarshal error:", err)
		return models.ListAuctionPageResponse{}, err
	}

	return res, nil
}
