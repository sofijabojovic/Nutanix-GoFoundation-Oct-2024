package main

import "fmt"

func main() {
	// var nos []int
	// var nos []int = []int{3, 1, 4, 2, 5}
	// var nos []int = []int{3, 1, 4}

	// type inference
	// var nos = []int{3, 1, 4, 2, 5}

	// using :=
	nos := []int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	// len()
	fmt.Println("len(nos) =", len(nos))

	// iterating using indexer
	fmt.Println("iterating using indexer")
	for idx := 0; idx < len(nos); idx++ {
		no := nos[idx]
		fmt.Printf("nos[%d] = %d\n", idx, no)
	}

	// iterating using range
	fmt.Println("iterating using range")
	for idx, no := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, no)
	}

	// append()
	fmt.Println("appending new values")
	nos = append(nos, 10)
	nos = append(nos, 20, 30, 40)

	hundreds := []int{100, 200, 300}
	nos = append(nos, hundreds...)
	fmt.Println(nos)

	// slicing
	fmt.Println("slicing")
	fmt.Println("nos[3:6] (idx 3 to 5) =", nos[3:6])
	fmt.Println("nos[3:] (idx 3 to end) =", nos[3:])
	fmt.Println("nos[:6] (idx 0 to 5) =", nos[:6])

	nos2 := nos
	nos[0] = 9999
	fmt.Printf("nos[0] = %d, nos2[0] = %d\n", nos[0], nos2[0])

}
