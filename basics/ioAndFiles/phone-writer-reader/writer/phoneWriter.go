package main

import "fmt"

type phoneWriter struct{}

func (p phoneWriter) Write(bs []byte) (int, error) {
	if len(bs) == 0 {
		return 0, nil
	}

	for i := 0; i < len(bs); i++ {
		if bs[i] >= '0' && bs[i] <= '9' {
			fmt.Print(string(bs[i]))
		}
	}

	fmt.Println()
	return len(bs), nil
}

func main() {
	phone1 := []byte("+1(234)567 90 10")
	phone2 := []byte("+2-345-678-12-35")

	writer := phoneWriter{}
	writer.Write(phone1)
	writer.Write(phone2)
}
