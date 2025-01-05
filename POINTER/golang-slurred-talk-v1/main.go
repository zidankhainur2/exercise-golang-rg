package main

import "fmt"

func SlurredTalk(words *string) {
	runes := []rune(*words)
	for i, r := range runes {
		switch r {
		case 'S', 'R', 'Z':
			runes[i] = 'L'
		case 's', 'r', 'z':
			runes[i] = 'l'
		}
	}
	*words = string(runes)
}

func main() {
	// bisa dicoba untuk pengujian test case
	var words string = "Saya Steven, saya suka menggoreng telur dan suka hewan zebra"
	SlurredTalk(&words)
	fmt.Println(words) // Output: Laya Lteven, laya luka menggoleng telul dan luka hewan lebra
}
