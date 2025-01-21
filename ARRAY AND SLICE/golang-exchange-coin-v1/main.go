package main

func ExchangeCoin(amount int) []int {
	if amount == 0 {
		return []int{}
	}

	coins := []int{1000, 500, 200, 100, 50, 20, 10, 5, 1}
	var result []int

	for _, coin := range coins {
		for amount >= coin {
			amount -= coin
			result = append(result, coin)
		}
	}

	return result
}
