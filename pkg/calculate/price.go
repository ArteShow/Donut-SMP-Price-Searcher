package calculate

func CalculateAveragePrice(items []float64) float64 {
	var sum float64
	for _, price := range items {
		sum += price
	}
	
	return sum / float64(len(items))
}