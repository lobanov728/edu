package main

import (
	"fmt"
	"log"
)

func main() {
	var res string
	res = longestCommonPrefix([]string{"flower", "flow", "flight"})
	if res != "fl" {
		fmt.Println(res)
		log.Fatalf("wrong")
	}

	res = longestCommonPrefix([]string{"dog", "racecar", "car"})
	if res != "" {
		fmt.Println(res)
		log.Fatalf("wrong")
	}
}
func longestCommonPrefix(strs []string) string {
	prefixes := make(map[string]int, 200)

	strsLen := len(strs)
	for i := range strs {
		for j := range strs[i] {
			index := strs[i][0 : j+1]
			prefixes[index]++
		}
	}
	l := 0
	longest := ""
	for i := range prefixes {
		if prefixes[i] == strsLen && len(i) > l {
			longest = string(i)
			l = len(i)
		}
	}

	return longest
}
