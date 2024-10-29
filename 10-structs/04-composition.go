package main

import "fmt"

type Product struct {
	id   int
	name string
	cost float64
}

// struct composition
type PerishableProduct struct {
	Product //composed
	expiry  string
	name    string
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
	/*
		var milk PerishableProduct
		milk.Product.id = 100
		milk.Product.name = "Milk"
		milk.Product.cost = 50
		milk.expiry = "2 Days"
	*/
	/*
		var milk PerishableProduct = PerishableProduct{
			Product: Product{
				id:   100,
				name: "Milk",
				cost: 50,
			},
			expiry: "2 Days",
			name:   "Flavoured Milk",
		}
	*/
	var milk *PerishableProduct = NewPerishableProduct(100, "Flavoured Milk", 50, "2 Days")
	fmt.Printf("%#v\n", milk)
	fmt.Println(milk.id, milk.name, milk.cost, milk.expiry)
}
