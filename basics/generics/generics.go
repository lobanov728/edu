package main

import "fmt"

// import "constraints"
type sss struct {
	in int32
}

func main() {
	x, y := 1, 3

	fmt.Println(fFirst(x, y))
	fmt.Println(fSSS(sss{1}, 2))
	fmt.Println(fComparable(x, y))
	fmt.Println(fMax(x, y))

	list := []int{1, 2, 4, 8}

	result := Reduce(list, func(a, b int) int {
		return a * b
	}, 1)
	fmt.Println("Reduce", result)

}

func fFirst[T any](x T, y T) T {
	return x
}

func fSSS[T any](x sss, y T) T {
	return y
}

func fComparable[T comparable](x, y T) T {
	if x == y {
		return x
	}

	return y
}

type ordered interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64
}

func fMax[T ordered](x, y T) T {
	if x > y {
		return x
	}

	return y
}

// func map(list, func) list // apply func to all element for the list
// func filter(list, func(el) bool) list //filter the list by result of the func

func Reduce[T any](list []T, accamulator func(T, T) T, init T) T {
	for _, el := range list {
		init = accamulator(init, el)
	}

	return init
}
