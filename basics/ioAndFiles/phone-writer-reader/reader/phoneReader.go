package main

import (
	"fmt"
	"io"
)

type phoneReader string

func (ph phoneReader) Read(p []byte) (int, error) {
	count := 0
	for i := 0; i < len(ph); i++ {
		if ph[i] >= '0' && ph[i] <= '9' {
			p[count] = ph[i]
			count++
		}
	}

	return count, io.EOF
}

func main() {
	phone1 := phoneReader("+1(234)567 90 10")
	phone2 := phoneReader("+2-345-678-12-35")

	buffer := make([]byte, len(phone1))
	phone1.Read(buffer)
	fmt.Println(buffer)
	fmt.Println(string(buffer))

	buffer2 := make([]byte, len(phone2))
	phone2.Read(buffer2)
	fmt.Println(buffer2)
	fmt.Println(string(buffer))
}
