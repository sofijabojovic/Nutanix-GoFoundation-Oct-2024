package main

import "fmt"

type Product struct {
	id   int
	name string
	cost float64
}

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.id, p.name, p.cost)
}

func (p *Product) ApplyDisount(discountPercentage float64) {
	p.cost = p.cost * ((100 - discountPercentage) / 100)
}

func main() {

	var pen Product = Product{
		id:   100,
		name: "Fountain Pen",
		cost: 20,
	}

	// fmt.Println(FormatProduct(pen))
	fmt.Println(pen.Format())

	fmt.Println("After applying 10% discount, pen ?")

	// ApplyDisount(&pen, 10)
	pen.ApplyDisount(10)

	// fmt.Println(FormatProduct(pen))
	fmt.Println(pen.Format())
}
