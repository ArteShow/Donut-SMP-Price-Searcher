package sort

import (
	"sort"

	"github.com/ArteShow/Donut-SMP-Price-Searcher/internal/models"
)

func SortArrayPricePerItem(array []models.AhObject) []models.AhObject {
	newArray := make([]models.AhObject, len(array))
	copy(newArray, array)

	sort.Slice(newArray, func(i, j int) bool {
		return (newArray[i].Price / newArray[i].Item.Count) <
			(newArray[j].Price / newArray[j].Item.Count)
	})

	return newArray
}
