package main

import (
	"fmt"
	"strings"
)

// hello World => d_l_r_o_W o_l_l_e_H
func ReverseString(str string) string {
	runes := []rune(str)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	var result strings.Builder
	for i, r := range runes {
		if r != ' ' {
			if i > 0 && runes[i-1] != ' ' {
				result.WriteRune('_')
			}
		}
		result.WriteRune(r)
	}
	return result.String()
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseString("Hello World"))
}
