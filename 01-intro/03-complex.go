package main

import "fmt"

func main() {
	var c1, c2 complex64
	c1 = 3 + 5i
	c2 = 7 + 3i
	fmt.Println(c1 + c2)
	fmt.Println("c1 = ", c1)
	fmt.Printf("real(c1) = %f\n", real(c1))
	fmt.Printf("imag(c1) = %f\n", imag(c1))
}
