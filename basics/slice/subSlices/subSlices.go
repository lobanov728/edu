package main

import "fmt"

func main() {
	//             0  1  2  3  4   5   6    7    8    9
	arr := [10]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512}
	s := []int{1, 2, 4, 8}

	slc := arr[3:6]
	fmt.Println(slc, len(slc), cap(slc)) // 8,16,32 len 3 cap 7
	c := append(slc, 1024)
	fmt.Println(c)
	fmt.Println(arr)

	slc = append([]int(nil), 10111)
	fmt.Println(slc, len(slc), cap(slc))
	fmt.Println(arr)

	slc1 := s[:1]
	fmt.Println(len(slc1), cap(slc1))
	slc2 := arr[1:2]
	fmt.Println(len(slc2), cap(slc2))
}
