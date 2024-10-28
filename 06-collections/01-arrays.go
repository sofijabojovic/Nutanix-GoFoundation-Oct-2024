package main

import "fmt"

func main() {
	// var nos [5]int
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}
	// var nos [5]int = [5]int{3, 1, 4}

	// type inference
	// var nos = [5]int{3, 1, 4, 2, 5}

	// using :=
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	// iterating using indexer
	fmt.Println("iterating using indexer")
	for idx := 0; idx < 5; idx++ {
		no := nos[idx]
		fmt.Printf("nos[%d] = %d\n", idx, no)
	}

	// iterating using range
	fmt.Println("iterating using range")
	for idx, no := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, no)
	}

	nos2 := nos // creates a copy coz arrays are values
	nos[0] = 9999
	fmt.Printf("nos[0] = %d, nos2[0] = %d\n", nos[0], nos2[0])

	fmt.Println("After sorting")
	sort(&nos)
	fmt.Println(nos)

	nosArrPtr := &nos
	fmt.Println("(*nosArrPtr)[0] = ", (*nosArrPtr)[0])
	fmt.Println("nosArrPtr[0] = ", nosArrPtr[0])

}

func sort(values *[5]int) {
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			/*
				if (*values)[i] > (*values)[j] {
					var temp = (*values)[i]
					(*values)[i] = (*values)[j]
					(*values)[j] = temp
				}
			*/
			/*
				if values[i] > values[j] {
					var temp = values[i]
					values[i] = values[j]
					values[j] = temp
				}
			*/

			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
}
