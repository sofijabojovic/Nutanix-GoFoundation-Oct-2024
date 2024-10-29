package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	for range 5 {
		wg.Add(1) // increment the counter by 1
		go f1(wg) // schedule the execution of f1 through the scheduler
	}
	f2()
	wg.Wait() // block until the counter becomes 0 (default = 0)
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done() // decrement the counter by 1
	fmt.Println("f1 started")
	// simulate time consuming operation
	time.Sleep(4 * time.Second)
	fmt.Println("f1 completed")
	return

}

func f2() {
	fmt.Println("f2 invoked")
}
