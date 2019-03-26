// Structs

// Struct is a type which contains named fields. These fields represents attributes. A struct is a typed collection of fields, useful to group data together to form records
// Structs are mutable
// You can also use pointers with structs. Pointers are automatically dereferenced.
package main

import (
	"fmt"
	"math"
)

// Circle is a new type created with the keyword `type`
type Circle struct {
	x float64
	y float64
	r float64
}

func main() {
	c := new(Circle) // var c Circle
	c1 := Circle{x: 0, y: 0, r: 5}
	c2 := Circle{1, 1, 2.5}

	// You can access fields using the `.` operator
	fmt.Println(c, c.x, c.y, c.r)
	fmt.Println(c1, c1.x)
	fmt.Println(c2, c2.r)

	c.x = 1
	c.y = 1
	c.r = 10
	fmt.Println(c, c.x, c.y, c.r)

	area := func(c Circle) float64 {
		return math.Pi * c.r * c.r
	}

	fmt.Println(area(c1))

	c0 := &c1
	c0.r = 2
	fmt.Println(c0.r, c1.r)
}
