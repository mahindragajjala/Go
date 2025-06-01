package closure

import "fmt"

/*
In Go, a closure is a function value that references variables from outside its body. The function can access and modify the variables even after the outer function has finished executing.

Think of a closure as "a function with its own environment" — it remembers the variables it was surrounded by.
*/
func ClosureFunction_outerfunc() {
	fmt.Println("OUTER FUNCTION...")
	inner := ClosureFunction_innerfunc()
	inner()
}
func ClosureFunction_innerfunc() func() {
	return func() {
		fmt.Println("INNER FUNCTION...")
	}
}
