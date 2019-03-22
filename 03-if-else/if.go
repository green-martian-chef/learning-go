/*
if/else-if/if
*/
package main

import "fmt"

func main() {
	// if/else if/else
	for n := 10; n > 0; n-- {
		if n%2 == 0 {
			// divisible by 2
			fmt.Println(n, "is even")
		} else {
			fmt.Println(n, "is odd")
		}
	}

	// FizzBuzz
	for n := 20; n > 0; n-- {
		if n%3 == 0 && n%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if n%5 == 0 {
			fmt.Println("Buzz")
		} else if n%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(n)
		}
	}
}
