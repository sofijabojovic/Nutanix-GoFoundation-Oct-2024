package main

import (
	"fmt"
	"log"
)

func main() {
	// ver 1.0
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		logAdd := getLoggedOperation(add)
		logSubtract := getLoggedOperation(subtract)

		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	add := getLoggedOperation(add)
	subtract := getLoggedOperation(subtract)

	add(100, 200)
	subtract(100, 200)

}

func getLoggedOperation(operationFn func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("Operation started")
		operationFn(x, y)
		log.Println("Operation completed")
	}
}

func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
