package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ChangeOutput(data []string) map[string][]string {
	temp := make(map[string]map[int][]string)

	for _, entry := range data {
		parts := strings.Split(entry, "-")
		if len(parts) != 4 {
			continue 
		}

		header := parts[0]
		index := parts[1]
		position := parts[2]
		value := parts[3]

		idx, err := strconv.Atoi(index)
		if err != nil {
			continue
		}

		if _, ok := temp[header]; !ok {
			temp[header] = make(map[int][]string)
		}

		if position == "first" {
			temp[header][idx] = append([]string{value}, temp[header][idx]...)
		} else if position == "last" {
			temp[header][idx] = append(temp[header][idx], value)
		}
	}

	result := make(map[string][]string)
	for header, indices := range temp {
		for i := 0; ; i++ {
			if vals, ok := indices[i]; ok {
				result[header] = append(result[header], strings.Join(vals, " "))
			} else {
				break
			}
		}
	}

	return result
}


// bisa digunakan untuk melakukan debug
func main() {
	data := []string{"account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe", "address-0-first-Jaksel", "address-0-last-Jakarta", "address-1-first-Bandung", "address-1-last-Jabar", "phone-0-first-081234567890", "phone-1-first-081234567891"}
	res := ChangeOutput(data)

	fmt.Println(res)
}
