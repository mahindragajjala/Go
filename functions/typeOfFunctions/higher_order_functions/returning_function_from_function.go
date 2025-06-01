package higherorderfunctions

import "fmt"

func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func FuncFromFunc() {
	double := multiplier(2)
	fmt.Println(double(5)) // Output: 10
}
