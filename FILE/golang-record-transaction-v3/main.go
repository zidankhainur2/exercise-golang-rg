package main

import (
	"fmt"
	"os"
	"sort"
)

type Transaction struct {
	Date   string
	Type   string
	Amount int
}

func RecordTransactions(path string, transactions []Transaction) error {
	// Map untuk menyimpan total income dan expense per tanggal
	dailyTotals := make(map[string]int)

	// Menghitung total income dan expense per tanggal
	for _, transaction := range transactions {
		if transaction.Type == "income" {
			dailyTotals[transaction.Date] += transaction.Amount
		} else if transaction.Type == "expense" {
			dailyTotals[transaction.Date] -= transaction.Amount
		}
	}

	// Mengurutkan tanggal
	dates := make([]string, 0, len(dailyTotals))
	for date := range dailyTotals {
		dates = append(dates, date)
	}
	sort.Strings(dates)

	// Membuka file untuk menulis
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Menulis hasil ke file
	for i, date := range dates {
		total := dailyTotals[date]
		transactionType := "income"
		if total < 0 {
			transactionType = "expense"
			total = -total
		}

		line := fmt.Sprintf("%s;%s;%d", date, transactionType, total)
		if i < len(dates)-1 {
			line += "\n"
		}

		_, err := file.WriteString(line)
		if err != nil {
			return err
		}
	}

	return nil
}


func main() {
	// bisa digunakan untuk pengujian test case
	var transactions = []Transaction{
		{"01/01/2021", "income", 100000},
		{"01/01/2021", "expense", 50000},
		{"01/01/2021", "expense", 30000},
		{"01/01/2021", "income", 20000},
	}

	err := RecordTransactions("transactions.txt", transactions)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success")
}
