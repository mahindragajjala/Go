package debugging

import "fmt"

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	Factorial(n - 1)
	fmt.Println(n)
	return n
}
func Factorial_main() {
	n := 5
	factorial := Factorial(n)
	fmt.Println(factorial)

}
