package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		d := RequestData(1)
		fmt.Println("res", d)
	}

	time.Sleep(time.Second)
	fmt.Println("runtime goroutine", runtime.NumGoroutine())
}

func RequestData(timeout time.Duration) string {
	dataChan := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		select {
		case dataChan <- requestFromSlowServer(ctx):
		case <-ctx.Done():
			close(dataChan)
			return
		}
	}()

	select {
	case res := <-dataChan:
		fmt.Printf("[+] request returned %s\n", res)
		return res
	case <-time.After(timeout):
		fmt.Println("[!] request timeout")
		cancel()
		return "empty val"
	}
}

func requestFromSlowServer(ctx context.Context) string {
	select {
	case <-time.After(time.Second):
		return "VIP Data"
	case <-ctx.Done():
		return "request interupted"
	}
}
