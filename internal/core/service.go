package core

import (
	"strconv"

	"github.com/ArteShow/Donut-SMP-Price-Searcher/internal/models"
	"github.com/ArteShow/Donut-SMP-Price-Searcher/pkg/avereg"
)

const GetAuctionItemsEndpoint = "/v1/auction/list/"

func CalculateAveregPrice(token, item string) (int64, error) {
	var responses []models.ListAuctionPageResponse
	counter := 1

	for {
		res, err := SendRequest(token, GetAuctionItemsEndpoint+strconv.Itoa(counter))
		if err != nil {
			break
		}
		if len(res.Response) == 0 {
			break
		}
		responses = append(responses, res)
		counter++
	}

	var prices []int64
	for _, res := range responses {
		for _, obj := range res.Response {
			if obj.Item.DisplayName == item {
				prices = append(prices, int64(obj.Price))
			}
		}
	}

	if len(prices) == 0 {
		return 0, nil
	}

	return avereg.GetAveregPrice(prices), nil
}
