package calc

func Sum(data ...int) int {
	total := 0
	for _, i := range data {
		total += i
	}
	return total
}
