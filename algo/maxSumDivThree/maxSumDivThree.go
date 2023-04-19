package main

import "fmt"

func main() {
	maxSumDivThree([]int{3, 6, 5, 1, 8})
}

func maxSumDivThree(nums []int) int {
	/*
	   опредление подзадачи
	   определение базы
	   вывод формулы перехода
	     [2,6,2,2,7] = 19 / 3   6  1 / 3
	                            5  4 / 3
	                            4  7 / 3
	                               0 1 2
	                            0      2
	                            1  6
	                            2      2
	                            3      2
	                            4    7
	     -----------------------------------------
	     3,6,5,1,8  =  23 / 3 = 7  2 / 3
	                            6   5 / 3
	                            5   8
	                            4   11
	                            3   14
	                            2   17
	                            1   20
	                            0   23
	*/
	// Represent the state as DP[pos][mod]: maximum possible sum
	// starting in the position "pos" in the array where the current
	// sum modulo 3 is equal to mod.

	// [2,6,2,2,7]
	/*
	  -> 2

	  -> 6

	  0 6
	  1 10
	  2 8
	*/

	dp := make([][3]int, len(nums))
	dp[0][nums[0]%3] = nums[0]

	for i := 1; i < len(nums); i++ {
		n := nums[i]
		dp[i] = dp[i-1]
		for _, acc := range dp[i-1] {
			mod := (n + acc) % 3
			dp[i][mod] = max(dp[i][mod], max(dp[i-1][mod], n+acc))
		}
	}
	fmt.Println(dp[len(nums)-1][0])

	return dp[len(nums)-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
