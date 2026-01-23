package core

import (
	"log"
	"strconv"
	"strings"

	"github.com/ArteShow/Donut-SMP-Price-Searcher/internal/models"
	"github.com/ArteShow/Donut-SMP-Price-Searcher/pkg/avereg"
)

const GetAuctionItemsEndpoint = "/v1/auction/list/"

func CalculateAveregPrice(token, item string) (float64, error) {
	log.Println("[CORE] calculating average price for:", item)

	var responses []models.ListAuctionPageResponse
	counter := 1

	for counter <= 20 {
		log.Println("[CORE] requesting page:", counter)

		res, err := SendRequest(token, GetAuctionItemsEndpoint+strconv.Itoa(counter))
		if err != nil {
			log.Println("[CORE] request failed:", err)
			break
		}

		log.Printf("[CORE] page %d items=%d\n", counter, len(res.Response))

		if len(res.Response) == 0 {
			log.Println("[CORE] empty page, stopping")
			break
		}

		responses = append(responses, res)
		counter++
	}

	var prices []float64
	for _, res := range responses {
		for _, obj := range res.Response {
			log.Printf("[CORE] checking item: %s | price=%d\n",
				obj.Item.DisplayName, obj.Price)

			if strings.Contains(
				strings.ToLower(obj.Item.DisplayName),
				strings.ToLower(item),
			) {
				log.Println("[CORE] MATCH FOUND")
				prices = append(prices, float64(obj.Price))
			}
		}
	}

	log.Println("[CORE] matched prices count:", len(prices))

	if len(prices) == 0 {
		log.Println("[CORE] no prices found")
		return 0, nil
	}

	avg := avereg.GetAveregPrice(prices)
	log.Println("[CORE] average price:", avg)

	return avg, nil
}
