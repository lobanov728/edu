package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("unable to open file", err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Println(file.Name())

}
