package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Readfile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(lines) == 0 {
		return []string{}, nil
	}

	return lines, nil
}

func CalculateProfitLoss(data []string) string {
	var totalIncome, totalExpense int
	var lastDate string

	for _, line := range data {
		parts := strings.Split(line, ";")
		if len(parts) != 3 {
			continue
		}

		date, transactionType, amountStr := parts[0], parts[1], parts[2]
		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			continue
		}

		lastDate = date
		if transactionType == "income" {
			totalIncome += amount
		} else if transactionType == "expense" {
			totalExpense += amount
		}
	}

	var resultType string
	var resultAmount int
	if totalIncome > totalExpense {
		resultType = "profit"
		resultAmount = totalIncome - totalExpense
	} else {
		resultType = "loss"
		resultAmount = totalExpense - totalIncome
	}

	return fmt.Sprintf("%s;%s;%d", lastDate, resultType, resultAmount)
}

func main() {
	// bisa digunakan untuk pengujian
	datas, err := Readfile("transactions.txt")
	if err != nil {
		panic(err)
	}

	result := CalculateProfitLoss(datas)
	fmt.Println(result)
}
