package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func DeliveryOrder(data []string, day string) map[string]float32 {
	adminFees := map[string]float32{
		"senin":  0.10,
		"rabu":   0.10,
		"jumat":  0.10,
		"selasa": 0.05,
		"kamis":  0.05,
		"sabtu":  0.05,
	}
	deliveryDays := map[string][]string{
		"senin":  {"JKT", "DPK"},
		"selasa": {"JKT", "BKS", "DPK"},
		"rabu":   {"JKT", "BDG"},
		"kamis":  {"JKT", "BDG", "BKS"},
		"jumat":  {"JKT", "BKS"},
		"sabtu":  {"JKT", "BDG"},
	}

	result := make(map[string]float32)

	validLocations := deliveryDays[day]

	for _, order := range data {
		parts := strings.Split(order, ":")
		if len(parts) != 4 {
			continue
		}

		firstName, lastName, priceStr, location := parts[0], parts[1], parts[2], parts[3]

		if !contains(validLocations, location) {
			continue
		}

		price, err := strconv.ParseFloat(priceStr, 32)
		if err != nil {
			continue
		}

		totalPrice := float32(price) * (1 + adminFees[day])
		totalPrice = float32(math.Round(float64(totalPrice)*100) / 100)
		key := firstName + "-" + lastName
		result[key] = totalPrice
	}

	return result
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func main() {
	data := []string{
		"Budi:Gunawan:10000:JKT",
		"Andi:Sukirman:20000:JKT",
		"Budi:Sukirman:30000:BDG",
		"Andi:Gunawan:40000:BKS",
		"Budi:Gunawan:50000:DPK",
	}

	day := "sabtu"

	deliveryData := DeliveryOrder(data, day)

	fmt.Println(deliveryData)
}
