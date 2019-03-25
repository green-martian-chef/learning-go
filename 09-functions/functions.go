/*
Functions

A function is an independent section of code that maps zero or more input parameters to zero or mor output parameters.


*/

package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0

	for _, v := range xs {
		total += v
	}

	return total / float64(len(xs))
}

// Returning multiple values
// Multiple values are often uset do return an error value along with the
// result (x, err := f()) or a boolean to indicate success (x, ok := f())
func multiple(a, b int) (int, int) {
	return a * b, b * b
}

// Named returns
func named() (r int) {
	r = 1
	return
}

// Variadic functions

// By using ... before the type name of the last parameter, you can indicate that it takes zero or more of those parameters.
func add(args ...int) int {
	var total int
	for _, v := range args {
		total += v
	}
	return total
}

// Closure
// A function inside another function where the inside function references  non-local variables form a closure. Another way to use closures is by writing a functions which returns another function.

var x int

func increment() int {
	x++
	return x
}

func makeEvenGenerator() func() int {
	i := 0
	return func() (num int) {
		num = i
		i += 2
		return
	}
}

// Recursion
// A function is able to call itself
func factorial(x uint) uint {
	/*
		x = 2
		x == 0 ? no -> x = 2
		then find factorial x - 1
			x == 0 ? no -> x = 1
				then find factorial x - 1
					x == 0 ? yes
					return 1
				return 1 * 1
		return 2 * 1
	*/
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

// Go has a statement called `defer` which schedules a function call to be run after the function completes. A deferred function is the last thing called by the program.
func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

// Examples
func half(n int) (int, bool) {
	if n%2 == 0 {
		return n / 2, true
	}
	return n / 2, false
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}

	fmt.Println(average(xs))

	fmt.Println(multiple(1, 2))

	fmt.Println(named())

	fmt.Println(add(1, 2, 3, 4, 5))

	fmt.Println("x", increment())
	fmt.Println("x", increment())
	fmt.Println("x", increment())

	nextEven := makeEvenGenerator() // returns a function which is assigned to var
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	fmt.Println(factorial(4))

	second()
	first()

	defer second()
	first()

	fmt.Println(half(1))
	fmt.Println(half(2))
	fmt.Println(half(3))
	fmt.Println(half(4))

	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(5))
	fmt.Println(fib(9))
}
