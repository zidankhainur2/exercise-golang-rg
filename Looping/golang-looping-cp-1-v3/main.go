package main

import "fmt"

func CountingNumber(n int) float64 {
	sum := 0.0
	current := 1.0

	for current <= float64(n) {
		sum += current
		current += 0.5
	}

	return sum
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingNumber(10))
}
