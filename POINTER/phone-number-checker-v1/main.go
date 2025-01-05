package main

import (
	"fmt"
	"strings"
)

func PhoneNumberChecker(number string, result *string) {
	if len(number) < 10 {
		*result = "invalid"
		return
	}

	if strings.HasPrefix(number, "62") {
		if len(number) < 11 {
			*result = "invalid"
			return
		}
		number = "0" + number[2:]
	} else if !strings.HasPrefix(number, "08") {
		*result = "invalid"
		return
	}

	prefix := number[:4]
	switch prefix {
	case "0811", "0812", "0813", "0814", "0815":
		*result = "Telkomsel"
	case "0816", "0817", "0818", "0819":
		*result = "Indosat"
	case "0821", "0822", "0823":
		*result = "XL"
	case "0827", "0828", "0829":
		*result = "Tri"
	case "0852", "0853":
		*result = "AS"
	case "0881", "0882", "0883", "0884", "0885", "0886", "0887", "0888":
		*result = "Smartfren"
	default:
		*result = "invalid"
	}
}

func main() {
	// bisa digunakan untuk pengujian test case
	var number = "081211111111"
	var result string

	PhoneNumberChecker(number, &result)
	fmt.Println(result)
}
