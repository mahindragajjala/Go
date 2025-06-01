package higherorderfunctions

import "fmt"

func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func add(x, y int) int {
	return x + y
}

func Passing_functions_as_arguments() {
	result := applyOperation(3, 4, add)
	fmt.Println("Result:", result) // Output: Result: 7
}
