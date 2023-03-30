package main

import "fmt"

type MyError struct{}

func (me MyError) Error() string {
	return "my error"
}
func main() {
	fmt.Println(returnError() == nil)          // true
	fmt.Println(returnErrorPtr() == nil)       // true
	fmt.Println(returnCustomError() == nil)    // false
	fmt.Println(returnCustomErrorPtr() == nil) // false
	fmt.Println(returnMyError() == nil)        // true
}

func returnMyError() *MyError {
	return nil
}

func returnCustomErrorPtr() error {
	var e *MyError
	return e
}

func returnCustomError() error {
	var e MyError
	return e
}

func returnErrorPtr() *error {
	var e *error
	return e
}

func returnError() error {
	var e error
	return e
}
