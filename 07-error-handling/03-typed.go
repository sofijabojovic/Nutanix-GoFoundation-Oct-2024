package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divisor cannot be zero")

func main() {
	if q, r, err := divide(100, 0); err != nil {
		if err == ErrDivideByZero {
			fmt.Println("Do not attempt to divide by zero")
			return
		}
		fmt.Println("UNKNOWN ERROR :", err)
	} else {
		fmt.Println(q, r)
	}
}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		err := errors.New("divisor cannot be zero")
		return 0, 0, err
	}
	quotient := x / y
	remainder := x % y
	return quotient, remainder, nil
}
*/

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = ErrDivideByZero
		return
	}
	quotient = x / y
	remainder = x % y
	return
}
