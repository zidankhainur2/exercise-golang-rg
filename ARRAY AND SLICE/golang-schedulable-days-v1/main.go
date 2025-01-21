package main

func SchedulableDays(date1 []int, date2 []int) []int {
	dateMap := make(map[int]bool)
	for _, date := range date1 {
		dateMap[date] = true
	}

	commonDates := []int{}
	for _, date := range date2 {
		if dateMap[date] {
			commonDates = append(commonDates, date)
		}
	}

	return commonDates
}
