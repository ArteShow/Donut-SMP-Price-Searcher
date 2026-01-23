package service

import (
	"log"

	"github.com/ArteShow/Donut-SMP-Price-Searcher/internal/requester"
)

func GetLowestAveragePricePerItem(item, token string) (float64, error) {
	page := 1
	items := 0
	var prices []float64
	for {
		if items < 20 {
			resp, err := requester.GetAhObjects(token, "lowest_price", item, page)
			if err != nil {
				break
			}

			log.Println(prices, resp)

		}else {
			break
		}
		
	}
}