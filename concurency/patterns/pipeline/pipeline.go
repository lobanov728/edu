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
	fmt.Println("addQuote")
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
	dataChan chan any
	pipe     []Pipe
}

func (p *pipeline) AddPipe(pipe Pipe) *pipeline {
	p.pipe = append(p.pipe, pipe)

	return p
}

func (p *pipeline) Execute() chan any {
	fmt.Println("Execute", len(p.pipe))
	for i := range p.pipe {
		outChan := make(chan any)
		inChan := p.dataChan
		go func() {
			fmt.Println("run pipe")
			defer close(outChan)
			for data := range inChan {
				fmt.Println("data from inChan", data)
				res, _ := p.pipe[i](data)
				outChan <- res
			}
		}()
		p.dataChan = outChan
	}

	return p.dataChan
}

func NewPipeline(filer func(chan any)) *pipeline {
	inputChan := make(chan any)

	go filer(inputChan)
	p := pipeline{
		dataChan: inputChan,
		pipe:     []Pipe{},
	}

	return &p
}

func main() {
	resultChan := NewPipeline(func(inputChan chan any) {
		defer close(inputChan)
		inputChan <- 104
		inputChan <- 10
	}).
		AddPipe(func(a any) (any, error) {
			return multiplyTwo(a.(int)), nil
		}).
		AddPipe(func(a any) (any, error) {
			return sqaure(a.(int)), nil
		}).
		AddPipe(func(a any) (any, error) {
			return addQuote(a.(int)), nil
		}).
		// AddPipe(func(a any) (any, error) {
		// 	return addFoo(a.(string)), nil
		// }).
		Execute()
	fmt.Println(len(resultChan))
	go func() {
		for {
			res, ok := <-resultChan
			if ok {
				fmt.Println(res)
			}
			break
		}
	}()

	doneChan := make(chan os.Signal, 1)
	signal.Notify(doneChan, syscall.SIGTERM)

	<-doneChan
	fmt.Println("finish")
}
