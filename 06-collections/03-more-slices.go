package main

import "fmt"

func main() {
	// var nos []int
	// var nos []int = make([]int, 0 /* initialize */, 3 /* memory allocation */)
	var nos []int = make([]int, 2 /* initialize */, 3 /* memory allocation */)
	fmt.Printf("len(nos) = %x, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 10)
	fmt.Printf("len(nos) = %x, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 20)
	fmt.Printf("len(nos) = %x, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 30)
	fmt.Printf("len(nos) = %x, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 40)
	fmt.Printf("len(nos) = %x, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 50)
	fmt.Printf("len(nos) = %x, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 60)
	fmt.Printf("len(nos) = %x, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 70)
	fmt.Printf("len(nos) = %x, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)

	nos2 := nos[4:]
	fmt.Printf("len(nos2) = %x, cap(nos2) = %d, nos2 = %v\n", len(nos2), cap(nos2), nos2)
}
