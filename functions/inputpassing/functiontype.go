package inputpassing

import "fmt"

/*
In Go, a "function type" means that
- functions themselves can be treated as values
— you can assign them to variables,
- pass them as arguments,
- or return them from other functions.

var f func(int, int) int  // f is a variable of function type

This means f is a variable that holds a function taking two int arguments and returning an int.
*/

// Define a function type
type Operation func(int, int) int

// Functions matching that type
func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

// A function that takes a function type as parameter
func compute(a, b int, op Operation) int {
	return op(a, b)
}

func Function_type() {
	var op Operation

	op = Add
	fmt.Println("Add:", compute(3, 4, op))

	op = Multiply
	fmt.Println("Multiply:", compute(3, 4, op))
}

/*
Real-World Uses of Function Types

Callbacks
Event handlers
Middleware (e.g., in HTTP servers)
Strategy pattern (choosing behavior at runtime)
Functional options pattern (as you’ve just seen)
*/
