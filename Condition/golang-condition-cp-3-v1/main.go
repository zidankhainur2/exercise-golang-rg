package main

import "fmt"

func GetPredicate(math, science, english, indonesia int) string {
	 hasil := (math + science + english + indonesia) / 4
	
	if  hasil == 100 {
		return "Sempurna"
	} else if hasil >= 90 {
		return "Sangat Baik"
	} else if hasil >= 80 {
		return "Baik"
	} else if hasil >= 70 {
		return "Cukup"
	} else if hasil >= 60 {
		return "Kurang"
	} else {
		return "Sangat kurang"
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetPredicate(50, 80, 100, 60))
}
