package main

import "fmt"

func main() {
	var i int = 1

	defer fmt.Println("result1 =>", func() int { return i * 2 }())
	defer func() { fmt.Println("result2 =>", func() int { return i * 2 }()) }()
	i++
	i++
	i++
	i++
	//выводит: result => 2 (not ok if you expected 4)
}
