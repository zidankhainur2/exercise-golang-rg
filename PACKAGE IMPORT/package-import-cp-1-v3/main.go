package main

import (
	"fmt"
	"strconv"
	"strings"
	"a21hc3NpZ25tZW50/internal"
)

func AdvanceCalculator(calculate string) float32 {
	if calculate == "" {
		return 0.0
	}

	tokens := strings.Split(calculate, " ")
	if len(tokens) == 0 {
		return 0.0
	}

	// Inisialisasi kalkulator dengan nilai awal
	initialValue, _ := strconv.ParseFloat(tokens[0], 32)
	calculator := internal.NewCalculator(float32(initialValue))

	// Iterasi melalui token dan lakukan operasi
	for i := 1; i < len(tokens); i += 2 {
		operator := tokens[i]
		value, _ := strconv.ParseFloat(tokens[i+1], 32)

		switch operator {
		case "+":
			calculator.Add(float32(value))
		case "-":
			calculator.Subtract(float32(value))
		case "*":
			calculator.Multiply(float32(value))
		case "/":
			calculator.Divide(float32(value))
		}
	}

	return calculator.Result()
}

func main() {
	res := AdvanceCalculator("3 * 4 / 2 + 10 - 5")
	fmt.Println(res)
}
