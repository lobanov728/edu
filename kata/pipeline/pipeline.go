package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func multiplyTwo(v int) int {
	fmt.Println("multiplyTwo")
	time.Sleep(time.Millisecond * 50)
	return v * 2
}

func sqaure(v int) int {
	fmt.Println("sqaure")
	time.Sleep(time.Millisecond * 50)
	return v * v
}

func addQuote(v int) string {
	fmt.Println("addQuote", v, fmt.Sprintf("'%d'", v))
	time.Sleep(time.Millisecond * 50)
	return fmt.Sprintf("'%d'", v)
}

func addFoo(v string) string {
	fmt.Println("addFoo")
	time.Sleep(time.Millisecond * 50)
	return fmt.Sprintf("foo - %s", v)
}

type Pipe func(any) (any, error)

type pipeline struct {
	inChan chan any
	pipes  []Pipe
}

func (p *pipeline) addPipe(pipe Pipe) *pipeline {
	p.pipes = append(p.pipes, pipe)

	return p
}

func (p *pipeline) Execute() (resultChan chan any) {
	for i := range p.pipes {
		i := i
		pipeOutChan := make(chan any)
		pipeInChan := p.inChan
		go func() {
			defer close(pipeOutChan)
			for input := range pipeInChan {
				res, _ := p.pipes[i](input)
				pipeOutChan <- res
			}
		}()
		p.inChan = pipeOutChan
	}

	return p.inChan
}

func NewPipeline() *pipeline {
	inputChan := make(chan any)

	go func() {
		defer close(inputChan)
		inputChan <- 3
	}()

	p := pipeline{
		inChan: inputChan,
		pipes:  []Pipe{},
	}

	return &p
}

func main() {
	some := NewPipeline().
		addPipe(func(a any) (any, error) {
			return multiplyTwo(a.(int)), nil
		}).
		addPipe(func(a any) (any, error) {
			return sqaure(a.(int)), nil
		}).
		addPipe(func(a any) (any, error) {
			return addQuote(a.(int)), nil
		}).
		addPipe(func(a any) (any, error) {
			return addFoo(a.(string)), nil
		}).
		Execute()

	for res := range some {
		fmt.Println("res", res)
	}

	doneChan := make(chan os.Signal)
	signal.Notify(doneChan, syscall.SIGINT)
	<-doneChan
	fmt.Println("finish")
}
