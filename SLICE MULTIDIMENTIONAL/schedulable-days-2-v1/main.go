package main

func SchedulableDays(villager [][]int) []int {
	if len(villager) == 0 {
		return []int{}
	}

	dateCount := make(map[int]int)

	for _, dates := range villager {
		for _, date := range dates {
			dateCount[date]++
		}
	}

	var result []int
	for date, count := range dateCount {
		if count == len(villager) {
			result = append(result, date)
		}
	}

	return result
}
