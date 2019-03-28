package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "This is a test string"

	// Return true if a string contains the substring
	fmt.Println(strings.Contains(str, "is"))

	// Count the number instances of a substring in a string
	fmt.Println(strings.Count(str, "t"))

	// Return the index of the first instance of the substring
	fmt.Println(strings.Index(str, "i"))

	// Concatenate the elements. Takes a separator
	fmt.Println(strings.Join([]string{"h", "i"}, ""))
	fmt.Println(strings.Join([]string{"first", "name"}, "-"))

	// Replace n instances of an old substring with a new substring
	fmt.Println(strings.Replace("numb3r", "3", "e", 0))
	fmt.Println(strings.Replace("numb3r", "3", "e", 1))
	fmt.Println(strings.Replace("numb33r", "3", "e", 1))
	fmt.Println(strings.Replace("numb33r", "3", "e", 2))

	str = "String"

	// Lowercase the string
	fmt.Println(str, strings.ToLower(str))

	// Uppercase the string
	fmt.Println(str, strings.ToUpper(str))
}
