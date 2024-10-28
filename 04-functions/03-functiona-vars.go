package main

import "fmt"

type OperationFn func(int, int) int

func main() {
	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi!")
	}
	sayHi()

	var divide func(int, int) (int, int)
	divide = func(x, y int) (quotient, remainder int) { //quotient & remainder are declared & initialized
		quotient, remainder = x/y, x%y
		return
	}
	q, r := divide(100, 7)
	fmt.Println(q, r)

	// var operation func(int, int) int
	var operation OperationFn
	operation = func(i1, i2 int) int {
		return i1 + i2
	}
	fmt.Println(operation(100, 200))

	operation = func(i1, i2 int) int {
		return i1 - i2
	}
	fmt.Println(operation(100, 200))

	operation = func(i1, i2 int) int {
		return i1 * i2
	}
	fmt.Println(operation(100, 200))

	operation = func(i1, i2 int) int {
		return i1 / i2
	}
	fmt.Println(operation(100, 200))
}
