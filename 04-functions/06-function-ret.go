package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var fn func()
	for range 10 {
		fn = getFn()
		fn()
		time.Sleep(500 * time.Millisecond)
	}
}

func getFn() func() {
	randNo := rand.Intn(100)
	switch {
	case randNo%2 == 0:
		return f1
	case randNo%3 == 0:
		return f2
	default:
		return func() {
			fmt.Println("anonymous fn invoked")
		}
	}
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
