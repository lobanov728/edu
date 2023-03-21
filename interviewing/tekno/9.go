package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	timeout := 3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	for {
		select {
		case <-time.After(time.Second):
			time.Sleep(5 * time.Millisecond)
			fmt.Println(1)
		case <-time.After(time.Second * 2):
			fmt.Println(2)
			cancel()
		case <-time.After(time.Second * 3):
			fmt.Println(3)
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		}

	}
}
