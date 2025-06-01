package typesOfRecursion

import "fmt"

/*
Linear recursion is a type of recursion where a function makes only one recursive call each time it runs.

Key Characteristics:
The problem is broken into one smaller subproblem.
Each call depends on a single recursive path.
The call stack forms a straight line (not a tree or graph).
*/

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1) // only one recursive call => linear recursion
}

func Linear_recursion() {
	fmt.Println("Factorial of 5:", factorial(5))
}
