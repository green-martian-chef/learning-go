// Goroutines
// Goroutine is a lightweight thread of execution

package main

import "fmt"

func f(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, ":", i)
	}
}

func main() {
	// f is called synchronously
	f("direct")

	// A goroutine starts with the keyword `go` and is executed concurrently with
	// the calling one
	go f("goroutine")

	// Both goroutines run asynchronously
	for i := 0; i < 3; i++ {
		// It is also possible to start a goroutine for an anonymous function call
		go func(m string, n int) {
			fmt.Println(m, ":", n)
		}("anonymous goroutine", i)
	}

	fmt.Scanln() // Requires a key to be pressed before the program exits
	fmt.Println("done")
}
