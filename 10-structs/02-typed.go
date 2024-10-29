package main

import "fmt"

type Product struct {
	id   int
	name string
	cost float64
}

func main() {

	var pen Product = Product{
		id:   100,
		name: "Fountain Pen",
		cost: 20,
	}
	// fmt.Printf("%+v\n", pen)
	fmt.Println(FormatProduct(pen))

	// assignment will result in a copy being created
	var p2 Product = pen
	p2.cost = 100
	fmt.Printf("pen.cost = %v, p2.cost = %v\n", pen.cost, p2.cost)

	var penPtr *Product
	penPtr = &pen
	// members can be directly accessed using "." notation from a pointer without deferencing
	fmt.Println(penPtr.id, penPtr.name, penPtr.cost)

	fmt.Println("After applying 10% discount, pen ?")
	ApplyDisount(&pen, 10)
	fmt.Println(FormatProduct(pen))
}

func FormatProduct(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.id, p.name, p.cost)
}

func ApplyDisount(p *Product, discountPercentage float64) {
	p.cost = p.cost * ((100 - discountPercentage) / 100)
}
