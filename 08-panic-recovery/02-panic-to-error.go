package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divide by zero error")

func main() {

	var divisor int = 0
	for {
		fmt.Println("Enter the divisor")
		fmt.Scanln(&divisor)
		fmt.Printf("[main] Attempting to divide 100 by %d\n", divisor)
		q, r, err := divideAdapter(100, divisor)
		if err != nil {
			fmt.Println("invalid input... try again!")
			continue
		}
		fmt.Printf("[main] quotient = %d and remainder = %d\n", q, r)
		break
	}

}

func divideAdapter(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

// 3rd party api
func divide(x, y int) (quotient, remainder int) {
	fmt.Println("[divide] calculating quotient")
	if y == 0 {
		panic(ErrDivideByZero)
	}
	quotient = x / y // results in a panic if y == 0
	fmt.Println("[divide] calculating remainder")
	remainder = x % y
	return
}
