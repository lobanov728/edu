package main

import (
	"fmt"
	"time"
)

type semaphore struct {
	semC chan struct{}
}

func NewSemaphore(maxConcurency int) semaphore {
	return semaphore{
		semC: make(chan struct{}, maxConcurency),
	}
}

func (s *semaphore) Acquire() {
	s.semC <- struct{}{}
}

func (s *semaphore) Release() {
	<-s.semC
}

func main() {
	sem := NewSemaphore(3)
	doneC := make(chan bool, 1)
	totalProcess := 10

	for i := 0; i < totalProcess; i++ {
		sem.Acquire()
		v := i
		go func() {
			defer sem.Release()
			fmt.Println(
				time.Now().Format("15:04:05"),
				"Running task with ID",
				v)
			time.Sleep(2 * time.Second)
			if v == totalProcess {
				doneC <- true
			}

		}()
	}
	<-doneC

}
