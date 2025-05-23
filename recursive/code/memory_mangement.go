package code

import "fmt"

// Recursive prints the value of `a` during each recursive call
// and stops recursion when a <= 0 (base case).
func Recursive(a int) int {
	if a <= 0 { // Base case
		fmt.Println("Reached base case with a =", a)
		return 0
	} else { // Recursive case
		fmt.Println("Calling Recursive with a =", a)
		return Recursive(a - 1)
	}
}

func Recursive_main() {
	Recursive(5)
}
