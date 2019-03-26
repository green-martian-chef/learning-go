// Pointers
//
// Pointers reference a location in memory where a value is stored rather than the value itself.
// A pointer is represented using `*` followed by the type of the stored value.
// `*` is also used to dereference pointer variables. Dereferencing a pointer gives access to the stored value instead of the memory address.
// `&` is used to find the memory address of the variable.
// It is also possible to create a pointer with the built-in `new()`
// Assigning a value to a dereferenced pointer changes the value at the referenced address
package main

import "fmt"

func zero(x int) {
	x = 0
}

func one(o *int) {
	*o = 1
}

// example
func square(x *float64) {
	*x = *x * *x
}

func main() {
	// In this example, x is not modified by zero() because when we call a
	// function that takes an argument, that argument is copied to the function
	x := 5
	y := &x
	o := 10

	zero(x)
	fmt.Println(x)
	fmt.Println(&x) // memory address
	fmt.Println(y)  // memory address, points to x
	fmt.Println(*y) // retrieves the value of x
	fmt.Println(o)  // original value
	one(&o)
	fmt.Println(o) // modified value

	p := new(int)
	fmt.Println(*p) // value 0
	fmt.Println(&p) // memory address
	*p = 5          // assign 5 to p
	fmt.Println(*p) // value 5

	v := 2.0
	square(&v)
	fmt.Println(v)

	a := 10
	b := &a
	*b = 20
	fmt.Println(a, *b)

	one := 1
	two := 2

	swap(&one, &two)

	fmt.Println(one, two)
}

func swap(x, y *int) {
	*x, *y = *y, *x
}
