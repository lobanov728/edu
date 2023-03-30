package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}

func twoSum(nums []int, target int) []int {
	storage := make(map[int]int, len(nums))
	for i, n := range nums {
		need := target - n
		stored, ok := storage[need]
		if ok {
			return []int{i, stored}
		}

		storage[n] = i
	}

	return []int{}
}
