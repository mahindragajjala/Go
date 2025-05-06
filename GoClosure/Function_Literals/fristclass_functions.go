package functionliterals

import "fmt"

/*
Go allows functions to be:

Assigned to variables
Passed as arguments
Returned from other functions
*/
//Passing functions as arguments
func operate(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(x, y int) int {
	return x + y
}

func Passing_functions_as_arguments() {
	result := operate(3, 4, add)
	fmt.Println("Result:", result) // Output: 7
}
