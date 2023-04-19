package main

import "fmt"

func main() {
	fmt.Println(countBits(12))
}

func countBits(n int) []int {
	//0  => 0000 0000 -  0 bits
	//1  => 0000 0001 -  1 bits

	//2  => 0000 0010 -  1 bits

	//3  => 0000 0011 -  2 bits

	//4  => 0000 0100 -  1 bits

	//5  => 0000 0101 -  2 bits
	//6  => 0000 0110 -  2 bits
	//7  => 0000 0111 -  3 bits

	//8  => 0000 1000 -  1 bits

	//9  => 0000 1001 -  2 bits
	//10 => 0000 1010  - 2 bits
	//11 => 0000 1011  - 3 bits
	//12 => 0000 1100  - 2 bits
	//13 => 0000 1101  - 3 bits
	//14 => 0000 1110  - 3 bits
	//15 => 0000 1111  - 4 bits

	//16 => 0001 0000  - 1 bits
	ones := make([]int, n+1)
	for x := 0; x <= n; x++ {
		fmt.Println(x, x&1, x>>1)
		ones[x] = ones[x>>1] + x&1
	}
	return ones
}
