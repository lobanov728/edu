package main

import "fmt"

type data struct {
	num   int
	key   *string
	items map[string]bool
}

func (this *data) pmethod() {
	this.num = 7 // will work
}

func (this data) vmethod() {
	this.num = 8                 // won't work
	(&this).num = 100            // won't work
	*(this.key) = "v.key1"       // will work
	this.items["vmethod"] = true // will work
}

func main() {
	key := "key.1"
	d := data{1, &key, make(map[string]bool)}

	fmt.Printf("num=%v key=%v items=%v\n", d.num, *d.key, d.items)
	// prints num=1 key=key.1 items=map[]

	d.pmethod()
	fmt.Printf("num=%v key=%v items=%v\n", d.num, *d.key, d.items)
	// prints num=7 key=key.1 items=map[]

	d.vmethod()
	fmt.Printf("num=%v key=%v items=%v\n", d.num, *d.key, d.items)
	// prints num=7 key=v.key items=map[vmethod:true]
}
