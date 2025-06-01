package code

import "fmt"

// Functions that take other functions as arguments or return them.
//syntax
func Apply(fn func(int, int) int, a, b int) int {
	return fn(a, b)
}

// Simple add function
func add(x int, y int) int {
	return x + y
}
func multiply(x int, y int) int {
	return x * y
}
func Higher_Order_Functions_main() {
	result1 := Apply(add, 10, 5)      // Calls add(10, 5)
	result2 := Apply(multiply, 10, 5) // Calls multiply(10, 5)

	fmt.Println("Sum:", result1)
	fmt.Println("Product:", result2)
}
