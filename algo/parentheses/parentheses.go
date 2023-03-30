package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("(){}[]"))
	fmt.Println(isValid("{((){}[])}"))
	fmt.Println(isValid("{((){}[}"))
	fmt.Println(isValid("){"))
	fmt.Println(isValid("[([]])"))
	fmt.Println(isValid("(("))
	fmt.Println(isValid("){"))
}

var closeBraceList map[rune]rune = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	var lastFromOrder rune
	braceOrder := make([]rune, 0, len(s)/2)

	for _, char := range s {
		closeBrace, ok := closeBraceList[char]
		if ok {
			braceOrder = append(braceOrder, closeBrace)
		} else {
			if len(braceOrder) == 0 {
				return false
			}

			lastFromOrder, braceOrder = braceOrder[len(braceOrder)-1], braceOrder[:len(braceOrder)-1]
			if lastFromOrder != char {
				return false
			}
		}
	}

	return len(braceOrder) == 0
}
