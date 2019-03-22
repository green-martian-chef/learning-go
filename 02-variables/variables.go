/*
Variables

A variable is a storage location, with a specific type and an associated name.package variables

Variables in Go are created by using the keywords `var` and `const`.
In Go, variables need to be declared before used.

To create a variable withou an assigned value:

keyword name type
------- ---- ----
var 		x 		int

To create a variable and assign a value to it:

var x int = 10

To create a variable and assign a value to it with no type specified:

x := 10

Go compiler is able to infer the type based on the literal value you assigned. The same is applied to the keyword `var`:

var x = 10

Generally, people use the short version `name := value`

Go is scoped using blocks. Variables exists within the block `{ }` they were declared but not outside it.


*/

package main

import "fmt"

func main() {
	// Assign string literal to x
	var x string
	x = "Hello, Variables"
	fmt.Println(x)

	// Assign string literal to y
	var y string
	y = "This is a variable"
	fmt.Println(y)

	// Re-assign a string literal to y
	y = "This variables has changed"
	fmt.Println(y)

	// Declare multiple variables at once
	var i, j int
	i = 0
	j = 1

	fmt.Println("i:", i)
	fmt.Println("j:", j)

	i += j
	fmt.Println("i:", i)
	i += j
	fmt.Println("i:", i)

	var (
		n = 10
		m = 20.0
	)

	fmt.Printf("n: %v, m: %v\n", n, m)

	// Scope
	outside()
	fmt.Println(o)

	// Example
	fahrenheit := 150.0                        // float64
	celsius := fahrenheitToCelsius(fahrenheit) // assigns a function to a variable
	feets := 3                                 // int
	meters := feetToMeters(feets)

	fmt.Println(fahrenheit, "Fahrenheit degrees are", celsius, "Celsius degrees")
	fmt.Println(feets, "feets are", meters, "meters.")
}

var o = "Outside"

func outside() {
	fmt.Println(o)
}
