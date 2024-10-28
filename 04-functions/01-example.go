package main

import "fmt"

func main() {
	sayHi()
	sayHello("Magesh")
	greet("Magesh", "Kuppan")
	fmt.Printf("add(100,200) = %d\n", add(100, 200))

	// fmt.Println(divide(100, 7))
	/*
		q, r := divide(100, 7)
		fmt.Printf("divide(100, 7), quotient = %d and remainder = %d\n", q, r)
	*/

	// using only the quotient
	// q, r := divide(100, 7)
	// q := divide(100, 7)
	q, _ := divide(100, 7)
	fmt.Printf("divide(100, 7), quotient = %d \n", q)
}

// 0 arguments, 0 returns
func sayHi() {
	fmt.Println("Hi!")
}

// 1 argument, 0 returns
func sayHello(name string) {
	fmt.Printf("Hello %s\n", name)
}

// 2 arguments, 0 returns
/*
func greet(firstName string, lastName string) {
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}
*/

// combining the parameters if they are of the same type
func greet(firstName, lastName string) {
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}

// 2 arguments, 1 return
func add(x, y int) int {
	return x + y
}

// 2 arguments, 2 returns
/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

// using named results
func divide(x, y int) (quotient, remainder int) { //quotient & remainder are declared & initialized
	/*
		quotient = x / y
		remainder = x % y
	*/
	quotient, remainder = x/y, x%y
	return
}
