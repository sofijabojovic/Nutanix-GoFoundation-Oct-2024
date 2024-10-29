package calculator

func Multiply(x, y int) int {
	opCount++
	// return x * y
	var result int
	for range x {
		result = Add(result, y) // "Add" function is directly accessible as both Multiply & Add belong to the same package (calculator)
	}
	return result
}
