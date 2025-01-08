package main

type Product struct {
	Name  string
	Price int
	Tax   int
}

func MoneyChanges(amount int, products []Product) []int {
	totalPrice := 0
	for _, product := range products {
		totalPrice += product.Price + product.Tax
	}
	change := amount - totalPrice
	if change <= 0 {
		return []int{}
	}
	denominations := []int{1000, 500, 200, 100, 50, 20, 10, 5, 1}
	result := []int{}
	for _, denom := range denominations {
		for change >= denom {
			change -= denom
			result = append(result, denom)
		}
	}

	return result
}
