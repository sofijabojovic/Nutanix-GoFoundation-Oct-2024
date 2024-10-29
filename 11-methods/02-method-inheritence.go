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

// struct composition
type PerishableProduct struct {
	Product //composed
	expiry  string
	name    string
}

// overriding Product.Format() method
func (pp PerishableProduct) Format() string {
	// return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f, Expiry = %q", pp.id, pp.name, pp.cost, pp.expiry)
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.expiry)
}

func NewPerishableProduct(id int, name string, cost float64, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			id:   id,
			cost: cost,
		},
		name:   name,
		expiry: expiry,
	}
}

func main() {

	var milk *PerishableProduct = NewPerishableProduct(100, "Flavoured Milk", 50, "2 Days")
	fmt.Printf("%#v\n", milk)
	milk.name = "Chocolate Milk"
	// fmt.Println(milk.id, milk.name, milk.cost, milk.expiry)
	fmt.Println(milk.Format())

	fmt.Println("After applying 10% discount")
	milk.ApplyDisount(10)
	fmt.Println(milk.Format())

}
