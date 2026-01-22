package avereg

func GetAveregPrice(Items []int64) int64 {
	var sum int
	for _, item := range Items {
		sum += int(item)
	}
	return int64(sum / len(Items))
}