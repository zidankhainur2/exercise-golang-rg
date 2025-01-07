package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ChangeOutput(data []string) map[string][]string {
	result := make(map[string][]string)

	for _, item := range data {
		parts := strings.Split(item, "-")
		if len(parts) != 4 {
			continue
		}

		header, index, position, value := parts[0], parts[1], parts[2], parts[3]

		if position == "first" {
			if len(result[header]) <= indexToInt(index) {
				result[header] = append(result[header], value)
			} else {
				result[header][indexToInt(index)] = value + " " + result[header][indexToInt(index)]
			}
		} else if position == "last" {
			if len(result[header]) <= indexToInt(index) {
				result[header] = append(result[header], "")
			}
			result[header][indexToInt(index)] = result[header][indexToInt(index)] + " " + value
		}
	}

	return result
}

func indexToInt(index string) int {
	i, _ := strconv.Atoi(index)
	return i
}

// bisa digunakan untuk melakukan debug
func main() {
	data := []string{"account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe", "address-0-first-Jaksel", "address-0-last-Jakarta", "address-1-first-Bandung", "address-1-last-Jabar", "phone-0-first-081234567890", "phone-1-first-081234567891"}
	res := ChangeOutput(data)

	fmt.Println(res)
}
