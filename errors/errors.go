package main

import (
	"errors"
	"fmt"
)

func myFunc(i int) (int, error) {
	if i == 42 {
		return 42, errors.New("you should never use 42")
	}

	return i, nil
}

type CustomError struct {
	code int
	msg  string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Custom error: code - %d: message - %s", e.code, e.msg)
}

func myFunc2(i int) (int, error) {
	if i == 42 {
		return 42, &CustomError{100, "you should never use 42"}
	}

	return i, nil
}

func main() {
	for _, v := range []int{12, 42} {
		if el, err := myFunc(v); err != nil {
			fmt.Println("Error: ", err.Error())
		} else {
			fmt.Println("Error: ", el)
		}
	}

	for _, v := range []int{12, 42} {
		if el, err := myFunc2(v); err != nil {
			fmt.Println("Error: ", err.Error())
		} else {
			fmt.Println("Error: ", el)
		}
	}

	fmt.Println("validate error type")
	_, e := myFunc2(42)
	if ae, ok := e.(*CustomError); ok {
		fmt.Println("Here:", ae.Error())
		fmt.Println(ae.code)
		fmt.Println(ae.msg)
	}
}
