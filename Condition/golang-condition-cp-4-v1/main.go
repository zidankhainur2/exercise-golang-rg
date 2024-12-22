package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
	
	tiket := VIP + regular + student
	VIP = VIP * 30
	regular = regular * 20
	student = student * 10


	if  (VIP + regular + student) <= 100 {
		return float32(VIP+regular+student)
	}

	if day%2 == 0 {
		if  tiket <= 5 {
			return float32(VIP+regular+student) * 0.9
		} else {
			return float32(VIP+regular+student) * 0.8
		}
	} else {
		if tiket <= 5 {
			return float32(VIP+regular+student) * 0.85
		} else {
			return float32(VIP+regular+student) * 0.75
		}
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(1, 1, 1, 20))
}
