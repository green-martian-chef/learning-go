// Methods
// Methods are a special function that add behavior to struct type.
// A method is just a function with a receiver before the name
// where the receiver is a parameter of the struc
// type Person struct {
//		name string
// }
// func (p *Person) talk() {
// 	 fmt.Println("Hi")
//}
// Methods can be defined for either pointer or value receiver types.
// Go automatically handles conversion between values and pointers for method calls.
package main

import (
	"fmt"
	"math"
)

// Circle represents a circle coordinates
type Circle struct {
	x, y, r float64
}

// Square represents a square coordinates
type Square struct {
	x, y float64
}

// Method
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Square) area() float64 {
	return c.x * c.y
}

func main() {
	circle := Circle{0, 0, 5}

	fmt.Println(circle.area())

	square := Square{2, 2}

	fmt.Println(square.area())
}
