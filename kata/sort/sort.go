package main

import (
	"fmt"
)

var iteration int

func quicksort(data []int) {
	iteration++
	fmt.Println("iteration", iteration)
	if len(data) < 2 {
		return
	}

	pivot := data[0]
	fmt.Println("pivot", pivot, len(data))

	l := 1
	r := len(data) - 1

	for r >= l {
		if data[l] > pivot {
			data[r], data[l] = data[l], data[r]
			r--
			continue
		}
		if data[l] <= pivot {
			l++
			continue
		}
	}

	data[l-1], data[0] = data[0], data[l-1]
	fmt.Println("data befor spliting", data)
	fmt.Println("data splited", data[:l-1], data[l:])
	quicksort(data[:l-1])
	quicksort(data[l:])
}

// 60, 54, 38, 87, 65, 21, 98, 74, 16, 59, 84, 65, 49, 31, 46, 98, 12
// 60, 54, 38, 87, 65, 40, 66
// 54 38 65 87 40 66
// 54 38 40 60 87 66 65
func main() {
	data := []int{
		60, 54, 38, 87, 65, 21, 98, 74, 16, 59, 84, 65, 49, 31, 46, 98, 12,
	}
	fmt.Println(len(data))

	quicksort(data)

	fmt.Println(data)
}
