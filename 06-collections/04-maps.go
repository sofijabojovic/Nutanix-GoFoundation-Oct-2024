package main

import "fmt"

func main() {
	/*
		var productRanks map[string]int = make(map[string]int)
		productRanks["pen"] = 5
		productRanks["pencil"] = 1
		productRanks["marker"] = 3
		productRanks["notepad"] = 4
	*/

	// var productRanks map[string]int = map[string]int{"pen": 5, "pencil": 1, "marker": 3}
	var productRanks map[string]int = map[string]int{
		"pen":     5,
		"pencil":  1,
		"marker":  3,
		"notepad": 4,
	}
	fmt.Println(productRanks)

	fmt.Println("len(productRanks) = ", len(productRanks))

	fmt.Println("Iterating a map")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	fmt.Println("Checking for the existence of a key")
	var nonExistentKey = "sharpner"
	if rank, exists := productRanks[nonExistentKey]; exists {
		fmt.Printf("productRanks[%q] = %d\n", nonExistentKey, rank)
	} else {
		fmt.Printf("key %q does not exist\n", nonExistentKey)
	}

	fmt.Println("Removing a key")
	keyToRemove := "sharpner"
	delete(productRanks, keyToRemove) // no-op if the key does not exist
	fmt.Println(productRanks)

	//
	dupProductRanks := productRanks // creates the copy of the POINTER
	productRanks["pen"] = 10
	fmt.Println("productRanks = ", productRanks)
	fmt.Println("dupProductRanks = ", dupProductRanks)
}
