package main

import (
	"fmt"
	"io"
)

type MyWriter struct {

}

func (m MyWriter) Write(p []byte) (n int, err error) {
	return
}


func main() {
	_, ok := interface{}(MyWriter{}).(io.Writer)
	fmt.Println(ok)
}

