package code

import "fmt"

//func(prefix string, values ...int)
func printSum(prefix string, values ...int) {
	sum := 0
	for _, v := range values {
		sum += v
	}
	fmt.Printf("%s %d\n", prefix, sum)
}

//Variadic_parameter
func Variadic_parameter() {
	printSum("Total:", 1, 2, 3, 4) // Output: Total: 10
}
