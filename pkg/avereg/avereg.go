package avereg

func GetAveregPrice(Items []float64) float64 {
	var sum int
	for _, item := range Items {
		sum += int(item)
	}
	return float64(sum / len(Items))
}