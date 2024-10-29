package main

import "fmt"

type Product struct {
	id   int
	name string
	cost float64
}

func main() {
	var p1 = Product{id: 100, name: "pen", cost: 10}
	var p2 = Product{id: 100, name: "pen", cost: 10}
	fmt.Println(p1 == p2)
}
