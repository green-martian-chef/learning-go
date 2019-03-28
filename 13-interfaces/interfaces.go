// Interfaces
// Interfaces are named collections of methods
// Interfaces define the behavior of a type
// In order to implement an interace, a type needs to implement the list of
// methods

package main

import "fmt"

type shape interface {
	area() float64
	perim() float64
}

// Since square implements the interface shape, it can be passed to functions
// that receive the interface shape as argument
type square struct {
	s float64
}

// square implements the interface shape. it has area() and perim() methods.
func (s square) area() float64 {
	return s.s * s.s
}

func (s square) perim() float64 {
	return s.s * 4
}

// A function can take interfaces as arguments
func measure(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perim())
}

func main() {
	s0 := square{s: 5}
	s1 := square{s: 2}
	measure(s0)
	measure(s1)
}
