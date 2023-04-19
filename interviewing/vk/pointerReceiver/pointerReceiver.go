package main

import "fmt"

type data struct {
	name string
}

func (p *data) print() {
	fmt.Println("name:", p.name)
}

func (p data) copy() {
	fmt.Println("copy:", p.name)
}

type printer interface {
	print()
	copy()
}

func check(p printer) {
	p.copy()
	p.print()
}

func main() {
	d1 := data{"one"}
	d1.print() //ok
	check(&d1)

	var in printer = data{"two"} // ошибка
	in.print()

	var in2 printer = &data{"two"} // ошибка
	in2.print()
	check(in2)

	m := map[string]data{"x": data{"three"}}
	m["x"].print() //ошибка
}
