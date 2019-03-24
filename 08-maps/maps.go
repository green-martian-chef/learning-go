/*
Maps

A map is an unordered collection of key-value pairs, also known as an associative array, hash table or dictionay.
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("maps")

	// the map type is represented by the keyword `map` followed by the key type in brackets and the value type
	// var x map[string]int // throw an error, maps need to be initialized before used
	x := make(map[string]int)

	arr := [5]string{
		"zero",
		"one",
		"two",
		"three",
		"four",
	}

	for i, v := range arr {
		x[v] = i
	}

	fmt.Println(x)
}
