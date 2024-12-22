package main

import "fmt"

func BMICalculator(gender string, height int) float64 {
	if condition := gender == "laki-laki"; condition {
		return float64(height - 100 ) - ((float64(height) - 100) * 0.1)
	} else {
		return float64(height - 100 ) - ((float64(height) - 100) * 0.15)
		
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BMICalculator("laki-laki", 165))
	fmt.Println(BMICalculator("perempuan", 165))
}
