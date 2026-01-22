package core

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/ArteShow/Donut-SMP-Price-Searcher/internal/models"
)

const BaseURL = "api.donutsmp.net"

func SendRequest(token, endpoint string) (models.ListAuctionPageResponse, error) {
	Request, err := json.Marshal(models.ListAuctionPageRequest{Sort: "lowest_price"})
	if err != nil {
		return models.ListAuctionPageResponse{}, err
	}
	
	req, err := http.NewRequest(http.MethodGet, BaseURL+endpoint, bytes.NewReader(Request))
	if err != nil {
		return models.ListAuctionPageResponse{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.ListAuctionPageResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.ListAuctionPageResponse{}, err
	}
	
	var res models.ListAuctionPageResponse
	if err = json.Unmarshal(body, &res); err != nil {
		return models.ListAuctionPageResponse{}, err
	}

	return res, nil
}