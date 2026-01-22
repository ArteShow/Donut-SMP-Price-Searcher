package avereg

func GetAveregPrice(Items []int64) int64 {
	var sum int64
	for _, item := range Items {
		sum += item
	}
	return sum / int64(len(Items))
}