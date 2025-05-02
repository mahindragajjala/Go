package typesOfRecursion

import "fmt"

// Function to check if a number is even
func isEven(n int) bool {
	if n == 0 {
		return true
	}
	return isOdd(n - 1)
}

// Function to check if a number is odd
func isOdd(n int) bool {
	if n == 0 {
		return false
	}
	return isEven(n - 1)
}

func Mutual_Recursion() {
	num := 7
	fmt.Printf("Is %d even? %v\n", num, isEven(num))
	fmt.Printf("Is %d odd? %v\n", num, isOdd(num))
}
