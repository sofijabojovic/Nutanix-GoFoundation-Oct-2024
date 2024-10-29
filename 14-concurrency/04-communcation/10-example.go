package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan string)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go genOdd(ch, wg)

	wg.Add(1)
	go genEven(ch, wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for no := range ch {
		fmt.Println(no)
	}
	fmt.Println("Done")
}

func genEven(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	no := 0
	for range 10 {
		time.Sleep(500 * time.Millisecond)
		ch <- fmt.Sprintf("Even : %d", no)
		no += 2

	}
}

func genOdd(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	no := 1
	for range 10 {
		time.Sleep(500 * time.Millisecond)
		ch <- fmt.Sprintf("Odd : %d", no)
		no += 2
	}
}
