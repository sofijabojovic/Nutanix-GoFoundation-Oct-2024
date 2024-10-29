package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divide by zero error")

func main() {
	defer func() {
		fmt.Println("[main] - deferred")
		if err := recover(); err != nil {
			fmt.Println("[main] - panic occurred, err :", err)
			return
		}
		fmt.Println("[main] - Thank you!")
	}()
	var divisor int = 0
	fmt.Printf("[main] Attempting to divide 100 by %d\n", divisor)
	q, r := divide(100, divisor)
	fmt.Printf("[main] quotient = %d and remainder = %d\n", q, r)
}

func divide(x, y int) (quotient, remainder int) {
	defer func() {
		fmt.Println("[divide] - deferred")
		if err := recover(); err != nil {
			fmt.Println("[divide] - panic occurred, err :", err)
			return
		}
	}()
	fmt.Println("[divide] calculating quotient")
	if y == 0 {
		panic(ErrDivideByZero)
	}
	quotient = x / y // results in a panic if y == 0
	fmt.Println("[divide] calculating remainder")
	remainder = x % y
	return
}
