package main

import (
	"context"
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"sync"
)
// GODEBUG=gctrace=1 ./{your app name}
// make request vie apache benchmark
// $ ab -k -c 8 -n 100000 "your url"
type someStruct struct {
	name string
	someSlice []int64
}

var sharedValue int64

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("start")
		ch := make(chan int64)
		go channelUpdater(ch)
		s := allocateAtStack()
		//s1 := allocateAtHeap()
		var sendOnly chan<- int64 = ch
		c := runBenchmark(sendOnly)

		for {

			if sharedValue > 1000 {
				fmt.Println(sharedValue)
				c()
				break
			}
		}

		writer.Write([]byte(strconv.Itoa(int(sharedValue))))
		writer.Write([]byte(s.name))

	})

	http.ListenAndServe(":3000", mux)
}

func allocateAtStack() someStruct {
	var s someStruct
	for i := 1; i < 100; i++ {
		s = someStruct{
			name:      "allocateAtStack",
			someSlice: make([]int64, 100),
		}
	}

	return s
}

func allocateAtHeap() *someStruct {
	var s someStruct
	for i := 1; i < 10; i++ {
		s = someStruct{
			name:      "allocateAtHeap",
			someSlice: make([]int64, 1000000),
		}
	}

	return &s
}

func channelUpdater(ch <-chan int64) {
	for x := range ch {
		sharedValue += x
	}
}

func runBenchmark(sendOnly chan<- int64) func() {
	sharedValue = 0
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(runtime.NumCPU())

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				select {
				case <-ctx.Done():
					wg.Done()
					return
				default:
					sendOnly <- 1
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(sendOnly)
	}()

	return cancel
}