package main

func main() {
	strStr("sadbutsad", "but")
}

func strStr(haystack string, needle string) int {
	result := -1
	needleLen := len(needle)
	haystackLen := len(haystack)
	br := true
	for i := 0; br; {
		end := i + needleLen
		if end > haystackLen {
			end = haystackLen
			br = false
		}
		sub := haystack[i:end]
		if sub == needle {
			return i
		}
		i++
	}

	return result
}
