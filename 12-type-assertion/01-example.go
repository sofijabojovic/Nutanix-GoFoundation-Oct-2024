package main

import "fmt"

type Product struct {
	id   int
	name string
	cost float64
}

func main() {
	var x interface{}
	x = 100
	x = "Veniam amet esse anim aute laborum commodo eu culpa consectetur occaecat anim in ad irure."
	x = 19.99
	x = true
	x = struct{}{}
	x = []int{10, 20, 30}
	fmt.Println(x)

	x = 100
	// x = "Dolor magna id id pariatur duis incididunt tempor commodo nulla."

	// unsafe
	// y := x.(int) + 200

	// safe
	if val, ok := x.(int); ok {
		y := val + 200
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	// type assertion using "type switch"
	// x = 100
	// x = "Veniam amet esse anim aute laborum commodo eu culpa consectetur occaecat anim in ad irure."
	x = Product{100, "Pen", 10}
	// x = 19.99
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x * 2 =", val*2)
	case string:
		fmt.Println("x is a string, len(x) = ", len(val))
	case Product:
		fmt.Println("x is a product :", val)
	default:
		fmt.Println("x is of unknown type")
	}

}
