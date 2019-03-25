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

	// Delete items from a map using built-in `delete()` and a key
	delete(x, "zero")
	fmt.Println(x)

	// A shorter way to creat a map
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	for e, n := range elements {
		fmt.Println(e, ":", n)
	}

	// The key does not exists, returns the zero value for the value type (empty string for type string)
	fmt.Println(elements["Un"])

	// Checks if the key exists: ok returns true or false
	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}
	if name, ok := elements["N"]; ok {
		fmt.Println(name, ok)
	}

	// Nested maps to store general information
	nestedElements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
	}

	if el, ok := nestedElements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
