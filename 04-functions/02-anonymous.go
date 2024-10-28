package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hi!")
	}()

	q, r := func(x, y int) (quotient, remainder int) { //quotient & remainder are declared & initialized
		quotient, remainder = x/y, x%y
		return
	}(100, 7)
	fmt.Println(q, r)
}
