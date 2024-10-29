package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // schedule the execution of f1 through the sche3duler
	f2()

	// blocking the main() so that the scheduler can look for other goroutines scheduled and execute them

	// poor man's synchronization techniques
	time.Sleep(500 * time.Millisecond)
	// fmt.Scanln()
}

func f1() {
	fmt.Println("f1 started")
	// simulate time consuming operation
	time.Sleep(2 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
