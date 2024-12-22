package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ReverseWord(str string) string {
	words := strings.Fields(str)
	for i, word := range words {
		runes := []rune(word)
		startUpper := unicode.IsUpper(runes[0])
		allUpper := true
		for _, r := range runes {
			if !unicode.IsUpper(r) {
				allUpper = false
				break
			}
		}
		for j, k := 0, len(runes)-1; j < k; j, k = j+1, k-1 {
			runes[j], runes[k] = runes[k], runes[j]
		}
		if startUpper && !allUpper {
			runes[0] = unicode.ToUpper(runes[0])
			for j := 1; j < len(runes); j++ {
				runes[j] = unicode.ToLower(runes[j])
			}
		}
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseWord("KITA AKAN SELALU BERSAMA"))
	fmt.Println(ReverseWord("A bird fly to the Sky"))
}
