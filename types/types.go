/*
Types

Data types categorize a set of related values, describe the operations that can be done on them and define the way they are stored.

Type Number

Go has several different types to represent numbers:
	* Integers: are numbers without a decimal component.
	* Floating point numbers: numbers that contain a decimal component.

Go's integer types are:
	* uint8
	* uint16
	* uint32
	* uint64
	* int8
	* int16
	* int32
	* int64

`uint` means 'unsigned integer' and int means 'signed integer'. `uint` only contain positive numbers and zero.

Go's floating points are:
	* float32
	* float64

	`float64` has increased precision.

Go also as two aditional types to represent complex numbers, `complex64` and `complex128`

Type String

String is a sequence of characters with a definite length used to represent text. Strings literals can be created using double quotes "Hello World" or back ticks `Hello World`. Double quotes allow special escape sequences as \n or \t.

Type Boolean

A boolean value is used to represent true or false. Three logical operators are used with boolean values:
	* && | and
	* || | or
	* !	 | not
*/
package main

import "fmt"

func main() {
	fmt.Println("Type Number\n")
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("1 + 1.0 =", 1+1.0)
	fmt.Println("1 + 1.5 =", 1+1.5)
	fmt.Println("5.0 - 2 =", 5.0-2)
	fmt.Println("4.5 - 2 =", 4.5-2)
	fmt.Println("4.5 - 2.0 =", 4.5-2.0)
	fmt.Println("5 / 2 =", 5/2)
	fmt.Println("5.0 / 2 =", 5.0/2)
	fmt.Println("5 % 2 =", 5%2)
	fmt.Println("")
	fmt.Println("Type Strings\n")
	fmt.Println("Hello String")
	fmt.Println(len("Hello String"))
	fmt.Println("Hello String"[1]) // returns the character `e` represented by a byte
	fmt.Println("Hello", "String") // performs concatenation
	fmt.Println("")
	fmt.Println("Type Boolean\n")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
