package main

import (
	"fmt"
	"time"
)

func main() {
	inch := make(chan int)
	outch := make(chan int)

	go func() {
		var in <-chan int = inch
		var out chan<- int
		var val int
		for {
			select {
			case out <- val:
				fmt.Println("1 case")
				out = nil
				in = inch
			case val = <-in:
				fmt.Println("2 case")
				out = outch
				in = nil
			}
		}
	}()

	go func() {
		for r := range outch {
			fmt.Println("result:", r)
		}
	}()

	time.Sleep(0)
	inch <- 1
	inch <- 2
	time.Sleep(3 * time.Second)
}
