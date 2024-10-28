package main

import "fmt"

func main() {

	var no int = 100
	var noPtr *int = &no // address of a value
	fmt.Println(noPtr)

	// deferencing - access the value from an address
	var val int = *noPtr
	fmt.Println(val)

	fmt.Println(no == *(&no))

	// pointers applied
	var x int = 100
	fmt.Println("[@main] &x = ", &x)
	fmt.Println("Before incrementing, x = ", x)
	increment(&x)
	fmt.Println("After incrementing, x = ", x)
}

func increment(z *int) {
	fmt.Println("[@increment] &z = ", z)
	*z += 1
}
