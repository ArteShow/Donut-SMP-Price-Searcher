package core

import (
	"strconv"

	"github.com/ArteShow/Donut-SMP-Price-Searcher/internal/models"
	"github.com/ArteShow/Donut-SMP-Price-Searcher/pkg/avereg"
)

const GetAuctionItemsEndpoint = "/v1/auction/list/"

func CalculateAveregPrice(token string) int64 {
	var Responses []models.ListAuctionPageResponse
	var counter int
	for {
		counter++
		res, err := SendRequest(token, GetAuctionItemsEndpoint+strconv.Itoa(int(counter)))
		if err != nil {
			break
		}
		Responses = append(Responses, res)
	}

	var Items []int64
	for _, res := range Responses {
		for _, item := range res.Response{
			Items = append(Items, int64(item.Price))
		}
	}

	return avereg.GetAveregPrice(Items)
}