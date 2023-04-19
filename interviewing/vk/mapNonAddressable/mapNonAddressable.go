package main

import "fmt"

type data struct {
	name string
}

func main() {
	m := map[string]data{"x": {"one"}}
	m["x"].name = "two" // error

	m1 := map[string]*data{"x": {"one"}}
	m1["x"].name = "two" //ok
	fmt.Println(m["x"])  //выводит: &{two}

	m2 := map[string]*data{"x": {"one"}}
	m2["z"].name = "what?" // runtime error: invalid memory address or nil pointer dereference
}
