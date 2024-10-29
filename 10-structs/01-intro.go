package main

import "fmt"

func main() {
	/*
		var pen struct {
			id   int
			name string
			cost float64
		}
		pen.id = 100
		pen.name = "Fountain Pen"
		pen.cost = 20
		// fmt.Println(pen)
		// fmt.Printf("%#v\n", pen)
		fmt.Printf("%+v\n", pen) */
	var pen struct {
		id   int
		name string
		cost float64
	} = struct {
		id   int
		name string
		cost float64
	}{
		id:   100,
		name: "Fountain Pen",
		cost: 20,
	}
	// fmt.Printf("%+v\n", pen)
	fmt.Println(FormatProduct(pen))
}

func FormatProduct(p struct {
	id   int
	name string
	cost float64
}) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.id, p.name, p.cost)
}
