package requester

import (
	byte "bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/ArteShow/Donut-SMP-Price-Searcher/internal/models"
)

const BaseURL = "https://api.donutsmp.net"
func GetAhObjects(token, sort, search string, page int) (models.GetAuctionStatsResponse, error) {
	body := models.GetAuctionStatsRequest{Sort: sort, Search: search}
	bytes, err := json.Marshal(body)
	if err != nil {
		return models.GetAuctionStatsResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, BaseURL+"/v1/auction/list/"+strconv.Itoa(page), byte.NewBuffer(bytes))
	if err != nil {
		return models.GetAuctionStatsResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.GetAuctionStatsResponse{}, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.GetAuctionStatsResponse{}, err
	}
	defer resp.Body.Close()

	var AhStats models.GetAuctionStatsResponse
	if err = json.Unmarshal(respBody, &AhStats); err != nil {
		return models.GetAuctionStatsResponse{}, err
	} 

	return AhStats, nil
}