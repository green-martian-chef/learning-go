/*
Arrays

An array is a numbered sequence of elements of a single type with a fixed length. Arrays in Go have a 0-based index.
*/
package main

import "fmt"

func main() {
	var x [5]int

	x[0] = 1

	fmt.Println(x)

	var y [5]float64
	y[0] = 92
	y[1] = 95
	y[2] = 84
	y[3] = 92
	y[4] = 96

	var total float64

	for i := 0; i < 5; i++ {
		total += y[i]
	}

	fmt.Println(total / 5)

	// Improved version
	total = 0

	for i := 0; i < len(y); i++ {
		total += y[i]
	}

	fmt.Println(total / float64(len(y))) //type conversiont. len(y) returns int

	/*
		The keyword `range`

		`range` iterates over elements in a variety of data structures.
		`range` on arrays and slices provides both the index and value for each entry.
	*/
	total = 0
	// A single _ tells the compiler we don't need the iterator variable
	for _, v := range y {
		total += v
	}

	fmt.Println(total / float64(len(y))) //type conversiont. len(y) returns int

	for i, v := range y {
		fmt.Println("Index:", i, "Value:", v)
	}

	// Shorter arrays
	z := [5]float64{
		100,
		99,
		98,
		97,
		96,
	}

	for i, v := range z {
		fmt.Println("Index, Value in Z:", i, v)
	}

	// The 4th element of an array
	fmt.Println("The 4h element of the array:", z[3])

	// Find the smallest number in the list
	n := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	min := n[0]

	for _, v := range n {
		if v < min {
			min = v
		}
	}

	fmt.Println(min)
}
