package main

import (
	"fmt"
	"strings"
)

func MyAtoi(s string) int {
	s = strings.Trim(s, " ")
	//for i, ch := range s {
	//	fmt.Println(s[i])
	//	fmt.Println(fmt.Sprintf("%T", s[i]))
	//	fmt.Println(fmt.Sprintf("%c", ch))
	//}
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}
	n := 0
	for _, ch := range []byte(s) {
		fmt.Println(ch)
		fmt.Println(byte('0'))
		ch -= '0'
		fmt.Println(ch)
		//fmt.Println(fmt.Sprintf("%c", ch))
		//fmt.Println(int(ch))
		//if ch > 9 {
		//	return 0, &NumError{fnAtoi, s0, ErrSyntax}
		//}
		//n = n*10 + int(ch)
	}
	return n
	//i, err := strconv.Atoi(s)
	//fmt.Println(err)
	//if i < math.MinInt32 {
	//	return math.MinInt32
	//}
	//if i > math.MaxInt32 {
	//	return math.MaxInt32
	//}
}

func main() {
	fmt.Println("String to integer casting")

}
