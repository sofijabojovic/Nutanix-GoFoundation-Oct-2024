package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for range 10 {
		time.Sleep(500 * time.Millisecond)
		if err := fn(); err != nil {
			fmt.Println("error occurred :", err)
			continue
		}
		fmt.Println("fn invocation successful")
	}
}

func fn() error {
	if rand.Intn(100)%3 == 0 {
		return errors.New("some random error")
	}
	return nil
}
