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
		log.Println("Operation started")
		add(100, 200)
		log.Println("Operation completed")

		log.Println("Operation started")
		subtract(100, 200)
		log.Println("Operation completed")
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	logOperation(add, 100, 200)
	logOperation(subtract, 100, 200)

}

func logOperation(operationFn func(int, int), x, y int) {
	log.Println("Operation started")
	operationFn(x, y)
	log.Println("Operation completed")
}

/*
func logAdd(x, y int) {
	log.Println("Operation started")
	add(x, y)
	log.Println("Operation completed")
}

func logSubtract(x, y int) {
	log.Println("Operation started")
	subtract(x, y)
	log.Println("Operation completed")
}
*/

func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
