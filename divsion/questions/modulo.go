package questions

import "fmt"

/*
Custom Modulo Function

Implement a function that performs modulo without using % operator.
*/

// CustomModulo returns the remainder of a divided by b without using the % operator
func CustomModulo(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}

	quotient := a / b
	remainder := a - (quotient * b)

	// Adjust remainder for negative numbers to mimic Go's % behavior
	if remainder < 0 && b > 0 || remainder > 0 && b < 0 {
		remainder += b
	}

	return remainder
}

func WithOut_Modulo() {
	fmt.Println(CustomModulo(10, 3))   // Output: 1
	fmt.Println(CustomModulo(-10, 3))  // Output: -1
	fmt.Println(CustomModulo(10, -3))  // Output: 1
	fmt.Println(CustomModulo(-10, -3)) // Output: -1
}
