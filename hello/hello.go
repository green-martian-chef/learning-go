/*
The anatomy of a Go program.

Go programs are read from top to bottom, left to right.

`package main` is a package declaration. All code in Go is organized in packages. There are two type of packages: executables and libraries.

The keyword `import` is how we include code from other packages into our package. Go has 25 reserved keywords.

`func main() { \\ code here }` declares a function named main. Functions are the building blocks of a Go program. All functions start with the keyword `func` followed by the name of the function. In Go, the `main` identifier is special. In Go, the execution of the program starts with the main function from the main package.
*/
package main

import "fmt"

func main() {
	/*
	  This statement prints on the screen the phrase 'Hello, world.'
	  Here, we use a function `Println()` from the package `fmt`.
	  To invoke a function from a package you need to write a statement following the bellow structure:

	  package         .  function()
	  -------------- --- ---------------
	  package's name dot function's name
	  -------------- --- ---------------
	  fmt.Println()
	  os.Exit()
	*/
	fmt.Println("Hello, world.")
}
