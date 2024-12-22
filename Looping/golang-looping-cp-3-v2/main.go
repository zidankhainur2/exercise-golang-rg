package main

import "fmt"

func CountingLetter(text string) int {
	unreadableLetters := "RSTZrstz"
	count := 0

	for _, char := range text {
		if contains(unreadableLetters, char) {
			count++
		}
	}

	return count
}

func contains(s string, c rune) bool {
	for _, char := range s {
		if char == c {
			return true
		}
	}
	return false
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Semangat"))
}
