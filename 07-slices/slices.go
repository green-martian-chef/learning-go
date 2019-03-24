/*
Slices

A slice is a segment of an array. It is indexed but unlike arrays, its length is not fixed.

This code `x := make([]float64, 5)` creates an slice that is associated with an underlying float64 array of length 5. Slices are always associated with some array and can never be longer than the array.

y := make([]float64, 5, 10) // 10 represents the capacity of the array, 5 represents the slice
*/

package main

import "fmt"

func main() {
	arr := [5]float64{1, 2, 3, 4, 5}

	x := arr[0:2]

	fmt.Println(x)

	y := arr[:2] // is the same as arr[0:2]
	z := arr[2:] // is the same as arr[2:5] or arr[2:len(arr)]

	fmt.Println(y, z)

	// Slices have 2 built-in functions `append` and `copy`
	arr1 := []int{1, 2, 3}

	// append() creates a new slice taking an existing slice and appending the following arguments
	arr2 := append(arr1, 4, 5)

	fmt.Println(arr1, arr2)

	// copy() copies the content of one slice to another
	arr3 := make([]int, 3)
	copy(arr3, arr2)
	fmt.Println(arr2, arr3)
}
