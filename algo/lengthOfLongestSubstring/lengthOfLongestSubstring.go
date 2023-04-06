package main

import (
	"fmt"
	"log"
)

func main() {
	// res := lengthOfLongestSubstring("a a b a a b ! b b")
	// if res != 3 {
	// 	log.Fatalf("wrong aabaab!bb, %d", res)
	// }
	res := lengthOfLongestSubstring("dzadc")
	if res != 4 {
		log.Fatalf("wrong dzadc, %d", res)
	}

	// res = lengthOfLongestSubstring("abcabcbb")
	// if res != 3 {
	// 	log.Fatalf("wrong abcabcbb")
	// }

	// res = lengthOfLongestSubstring("pwwkew")
	// if res != 3 {
	// 	log.Fatalf("wrong pwwkew %d", res)
	// }

	// res = lengthOfLongestSubstring("ckilbkd")
	// if res != 5 {
	// 	log.Fatalf("wrong pwwkckilbkdew %d", res)
	// }
	// res = lengthOfLongestSubstring("tmmzuxt")
	// if res != 5 {
	// 	log.Fatalf("wrong tmmzuxt %d", res)
	// }

	// res = lengthOfLongestSubstring("wobgrovw")
	// if res != 6 {
	// 	log.Fatalf("wrong wobgrovw %d", res)
	// }
}

func lengthOfLongestSubstring(s string) int {
	var dict [128]bool
	result, length := 0, 0

	// a a b a a b ! b b

	for i, j := 0, 0; i < len(s); i++ {
		index := s[i]
		fmt.Println("index", s[i], string(s[i]))
		if dict[index] {
			for ; dict[index]; j++ {
				fmt.Println("s[j]", s[j], string(s[j]))
				length--
				dict[s[j]] = false
			}
		}
		dict[index] = true
		length++
		if length > result {
			result = length
		}

	}

	return result
}
