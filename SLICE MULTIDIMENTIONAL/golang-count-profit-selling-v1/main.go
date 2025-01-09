package main

func CountProfit(data [][][2]int) []int {
	if len(data) == 0 {
		return []int{}
	}


	numMonths := len(data[0])
	totalProfits := make([]int, numMonths)

	for _, branch := range data {
		for monthIndex, salesAndExpenses := range branch {
			profit := salesAndExpenses[0] - salesAndExpenses[1]
			totalProfits[monthIndex] += profit
		}
	}

	return totalProfits
}
