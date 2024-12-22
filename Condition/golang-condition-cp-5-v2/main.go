package main

import "fmt"

func TicketPlayground(height, age int) int {
	
	total := 0
	if age < 5 {
		return -1
	} else if age >= 5 && age <=7 {
		if height > 135 && height <= 150 {
			total = 25000
		} else if height > 150 && height <= 160 {
			total = 40000
		} else if height > 160 {
			total = 60000
		} else if height >120 && height <= 135 {
			total = 15000
		}
	} else if age >= 8 && age <=9 {
		if height > 135 && height <= 150 {
			total = 25000
		} else if height > 150 && height <= 160 {
			total = 40000
		} else if height > 160 {
			total = 60000
		} else if height > 120 && height <= 135 {
			total = 25000
		}
	} else if age >= 10 && age <=11 {
		if height > 150 && height <= 160 {
			total = 40000
		} else if height > 160 {
			total = 60000
		} else if height > 120 && height <= 150 {
			total = 40000
		}
	} else if age == 12 {
		if height > 160 {
			total = 60000
		} else if height > 120 && height <= 160 {
			total = 60000
		}
	} else if age > 12 {
		total = 100000
	}
	
	return total
	
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(TicketPlayground(160, 11))
}
