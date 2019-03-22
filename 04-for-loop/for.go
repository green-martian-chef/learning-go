/*
For Loop

The `for` statement allows us to repeat a block multiple times.
Go doesn't have different types of loops like `while`, `until`, `do` or `foreach`.
*/
package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Inline for loop, the Go way
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// for loop with `range` /* don't */
	for i := range [10]int{} {
		fmt.Println(i + 1)
	}
}
