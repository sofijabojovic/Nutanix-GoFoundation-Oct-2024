package main

import (
	"errors"
	"fmt"
)

func main() {
	if q, r, err := divide(100, 0); err != nil {
		fmt.Println("ERROR :", err)
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
		err = errors.New("divisor cannot be zero")
		return
	}
	quotient = x / y
	remainder = x % y
	return
}
