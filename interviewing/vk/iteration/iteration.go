package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println("field p name", p.name)
}

func main() {
	data := []string{"one", "two", "three"}

	for _, v := range data {
		go func() {
			fmt.Println(v)
		}()
		// горутины выводят: three, three, three
	}

	for i := range data {
		i := i
		go func() {
			fmt.Println(data[i])
		}()
		// горутины выводят: one, two, three
	}

	for _, v := range data {
		vcopy := v
		go func() {
			fmt.Println(vcopy)
		}()
		// горутины выводят: one, two, three
	}

	data1 := []*field{{"one"}, {"two"}, {"three"}}

	for _, v := range data1 {
		go v.print()
		// горутины выводят: one, two, three
	}

	time.Sleep(3 * time.Second)

}
