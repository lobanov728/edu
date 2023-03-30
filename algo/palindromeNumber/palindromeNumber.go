package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
}

func isPalindrome(num int) bool {
	if num < 0 {
		return false
	}

	var reverted int
	x := num
	for x > 0 {
		reverted = reverted*10 + (x % 10)
		x = (x - x%10) / 10
	}

	return reverted == num
}
