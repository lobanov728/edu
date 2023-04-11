package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

type Ball struct {
	hits int
}

var wg sync.WaitGroup

func main() {
	ch := make(chan Ball)
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(2)
	go player(ctx, "Alice", ch)
	go player(ctx, "Bob", ch)

	ch <- Ball{}
	time.Sleep(time.Second)

	fmt.Println("before", runtime.NumGoroutine())
	cancel()
	<-ch
	close(ch)
	wg.Wait()
	fmt.Println("after", runtime.NumGoroutine())
}

func player(ctx context.Context, name string, ch chan Ball) {
	for {
		if ctx.Err() != nil {
			wg.Done()
			log.Printf("Player %s, stop playing", name)
			return
		}

		select {
		case <-ctx.Done():
			wg.Done()
			log.Printf("Player %s, stop playing", name)
			return
		case b := <-ch:
			fmt.Println("ctx.Err() != nil1", name, ctx.Err() != nil)
			fmt.Println("ctx.Err() != nil2", name, ctx.Err() != nil)
			b.hits++
			log.Printf("Player %s, ball hits %d %p", name, b.hits, &b)
			fmt.Println("ctx.Err() != nil3", name, ctx.Err() != nil)
			time.Sleep(time.Millisecond * 100)
			fmt.Println("ctx.Err() != nil4", name, ctx.Err() != nil)
			ch <- b
			fmt.Println("ctx.Err() != nil5", name, ctx.Err() != nil)
		}
	}
}
