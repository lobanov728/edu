package main

import "fmt"

func main() {
	var numbers []*int

	for _, val := range []int{10, 20, 30, 40} {
		numbers = append(numbers, &val)
		fmt.Println(&val)
	}

	for _, number := range numbers {
		fmt.Printf("%d ", *number)
	}
}
