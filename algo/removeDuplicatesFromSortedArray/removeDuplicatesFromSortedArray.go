package main

func main() {
	// result 0 1 2 3 4 and 5
	removeDuplicates([]int{0, 0, 1, 1, 2, 3, 4, 4})

}

func removeDuplicates(nums []int) int {
	ln := len(nums)
	if ln <= 1 {
		return ln
	}

	j := 0
	for i := 1; i < ln; i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}

	return j + 1
}
