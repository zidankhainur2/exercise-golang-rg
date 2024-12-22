package main

import (
	"fmt"
	"strconv"
)


func BiggestPairNumber(numbers int) int {
	strNum := strconv.Itoa(numbers)
	maxPair := 0
	maxSum := 0

	for i := 0; i < len(strNum)-1; i++ {
		pair, _ := strconv.Atoi(strNum[i : i+2])

		sum := int(strNum[i]-'0') + int(strNum[i+1]-'0')

		if sum > maxSum || (sum == maxSum && pair > maxPair) {
			maxPair = pair
			maxSum = sum
		}
	}

	return maxPair
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BiggestPairNumber(11223344))
}
