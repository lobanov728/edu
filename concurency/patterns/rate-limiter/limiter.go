package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type limiter struct {
	C chan struct{}
}

func NewLimiter(concurencyLvl int64, rate time.Duration) limiter {
	l := limiter{
		C: make(chan struct{}, concurencyLvl),
	}

	go func() {
		for {
			tick := time.NewTicker(rate)
			for i := 0; i < int(concurencyLvl); i++ {
				l.C <- struct{}{}
			}
			<-tick.C
		}
	}()

	return l
}

// add limiter with limit 10 per second
// read values concurrently
func main() {

	count := 50
	ch := make(chan int, count)
	limiter := NewLimiter(10, time.Second)

	wg := sync.WaitGroup{}
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-limiter.C
			ch <- RPCCall()
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < count; i++ {
			fmt.Println(<-ch)
		}
	}()
	wg.Wait()
}

func RPCCall() int {
	return rand.Int()
}
