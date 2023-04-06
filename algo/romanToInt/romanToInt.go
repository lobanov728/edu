package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIIV"))
}

func romanToInt(s string) int {
	matches := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	total := 0
	var sign int
	b := []byte(s)
	prev := 0
	for i := len(b) - 1; i >= 0; i-- {
		num := matches[b[i]]

		if num < prev {
			sign = -1
		}
		if num > prev {
			sign = 1
		}
		prev = num
		total += num * sign
	}

	return total
}
