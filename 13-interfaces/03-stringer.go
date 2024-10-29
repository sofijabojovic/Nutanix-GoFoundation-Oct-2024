package main

import (
	"fmt"
	"math"
)

// ver 1.0
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Length  float64
	Breadth float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case Circle:
		fmt.Println("Area :", val.Area())
	case Rectangle:
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("x is neither a circle nor a rectangle")
	}
}
*/

/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case interface{ Area() float64 }:
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("x does not have Area() method")
	}
}
*/

type AreaFinder interface{ Area() float64 }

func PrintArea(x AreaFinder) {
	fmt.Println("Area :", x.Area())
}

/* ********** ver 1.0 ***********/

// ver 2.0
// Perimeter method for circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Perimeter method for rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Breadth)
}

type PerimeterFinder interface{ Perimeter() float64 }

func PrintPerimeter(x PerimeterFinder) {
	fmt.Println("Perimeter :", x.Perimeter())
}

// ver 3.0

type ShapeStatsFinder interface {
	AreaFinder
	PerimeterFinder
}

func PrintShapeStats(x ShapeStatsFinder) {
	fmt.Println("Shape Stats:")
	PrintArea(x)      // x should be interface{ Area() float64 }
	PrintPerimeter(x) // x should be interface { Perimeter() float64 }
}

// ver 4.0
// fmt.Stringer interface implementation
func (c Circle) String() string {
	return fmt.Sprintf("CIRCLE :: Radius = %0.2f, Area = %0.2f, Perimeter = %0.2f", c.Radius, c.Area(), c.Perimeter())
}

func (r Rectangle) String() string {
	return fmt.Sprintf("RECTANGLE :: Length = %0.2f, Breadth = %0.2f, Area = %0.2f, Perimeter = %0.2f", r.Length, r.Breadth, r.Area(), r.Perimeter())
}

func main() {
	c := Circle{Radius: 12}
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	// PrintShapeStats(c)
	fmt.Println(c)

	r := Rectangle{Length: 12, Breadth: 14}
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	// PrintShapeStats(r)
	fmt.Println(r)

}
