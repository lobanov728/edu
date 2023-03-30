package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 3)
	for i := 0; i < 5; i++ {
		go func() {
			time.Sleep(time.Second)
			ch <- i
			fmt.Println("after write")
		}()
	}

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(<-ch)
			fmt.Println("after read")
		}()
	}
	time.Sleep(time.Second * 10)
}
