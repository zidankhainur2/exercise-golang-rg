package main

func ReverseData(arr [5]int) [5]int {
	var result [5]int
	for i, v := range arr {
		reversedIndex := len(arr) - 1 - i
		reversedValue := reverseNumber(v)
		result[reversedIndex] = reversedValue
	}
	return result
}

func reverseNumber(n int) int {
	reversed := 0
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return reversed
}
