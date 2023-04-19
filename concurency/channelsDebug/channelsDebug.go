package main

import "time"

func main() {
	ch := make(chan int32, 4)

	ch <- 1
	ch <- 2

	_ = <-ch

	ch <- 3

	go func() {
		for {
			ch <- 42
		}
	}()
	time.Sleep(time.Second)

	return
}
