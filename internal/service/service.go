package service

import (
	"log"

	"github.com/ArteShow/Donut-SMP-Price-Searcher/internal/requester"
	"github.com/ArteShow/Donut-SMP-Price-Searcher/pkg/calculate"
	"github.com/ArteShow/Donut-SMP-Price-Searcher/pkg/sort"
)

func GetLowestAveragePricePerItem(item, token string, amount float64) (float64, error) {
	log.Printf("[START] item=%s amount=%.2f", item, amount)

	page := 1
	itemsCollected := 0.0
	prices := make([]float64, 0, int(amount))

	for itemsCollected < amount {
		log.Printf("[REQUEST] page=%d itemsCollected=%.2f", page, itemsCollected)

		resp, err := requester.GetAhObjects(token, "lowest_price", item, page)
		if err != nil {
			log.Printf("[ERROR] request failed on page %d: %v", page, err)
			break
		}

		if len(resp.Items) == 0 {
			log.Printf("[STOP] no items returned on page %d", page)
			break
		}

		log.Printf("[RESPONSE] page=%d items=%d", page, len(resp.Items))

		sorted := sort.SortArrayPricePerItem(resp.Items)
		log.Printf("[SORT] page=%d sorted by price per item", page)

		for idx, i := range sorted {
			perItemPrice := i.Price / i.Item.Count
			remaining := amount - itemsCollected
			stackCount := i.Item.Count

			take := stackCount
			if take > remaining {
				take = remaining
			}

			log.Printf(
				"[ITEM] #%d price=%.2f count=%.2f perItem=%.2f take=%.2f remaining=%.2f",
				idx,
				i.Price,
				i.Item.Count,
				perItemPrice,
				take,
				remaining,
			)

			for j := 0; j < int(take); j++ {
				prices = append(prices, perItemPrice)
				itemsCollected++
			}

			if itemsCollected >= amount {
				log.Printf("[ENOUGH] collected=%.2f needed=%.2f", itemsCollected, amount)
				break
			}
		}

		page++
	}

	if len(prices) == 0 {
		log.Printf("[RESULT] no prices collected → returning 0")
		return 0, nil
	}

	avg := calculate.CalculateAveragePrice(prices)
	log.Printf(
		"[RESULT] pricesCollected=%d average=%.2f",
		len(prices),
		avg,
	)

	return avg, nil
}
