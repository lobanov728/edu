package main

import "fmt"

func main() {
	test := map[int]int{1: 123}
	fmt.Printf("%#v, %v\n", test, test[123])
	if el, ok := test[1]; ok {
		fmt.Println(el)
	}



}
