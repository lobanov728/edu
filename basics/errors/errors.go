package main

import "fmt"

type MyError struct {
}

func (e MyError) Error() string {
	return "MyError string"
}

func main() {
	fmt.Println("returnError", returnError() == nil)                            // true
	fmt.Println("returnErrorPtr1", returnErrorPtr1(), returnErrorPtr1() == nil) // true
	fmt.Println("returnErrorPtr2", returnErrorPtr2(), returnErrorPtr2() == nil) // false
	fmt.Println("returnMyError", returnMyError() == nil)                        // false
	fmt.Println("returnMyErrorPtr1", returnMyErrorPtr1() == nil)                // false
	fmt.Println("returnMyErrorPtr2", returnMyErrorPtr2() == nil)                // false
	fmt.Println("work", work() == nil)
}

func returnError() error {
	var err error
	return err
}

func returnErrorPtr1() *error {
	var err *error
	return err
}

func returnErrorPtr2() *error {
	var err error
	return &err
}

func returnMyError() error {
	var err MyError
	return err
}

func returnMyErrorPtr1() error {
	var err *MyError
	return err
}

func returnMyErrorPtr2() error {
	var err MyError
	return &err
}

func work() *MyError {
	return nil
}
