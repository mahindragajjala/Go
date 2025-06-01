package code

import "fmt"

func Factorial(x int) int {
	if x == 1 {
		return 1
	}
	return x * Factorial(x-1)
}
func Factorial_main() {
	factorial := Factorial(5)
	fmt.Println(factorial)
}
