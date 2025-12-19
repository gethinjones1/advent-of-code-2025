package day6

func CalculateMath(sums [][]interface{}) int {
	var total int
	for i := 0; i < len(sums)-1; i++ {
		if sums[len(sums)-1][0] == "+" {
			total += sums[i][0].(int)
			continue
		}
		if total == 0 {
			total = 1
		}
		total *= sums[i][0].(int)
	}
	return total
}
