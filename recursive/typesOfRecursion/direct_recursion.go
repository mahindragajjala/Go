package typesOfRecursion

import "fmt"

func Direct_recursion() {
	printNum(5)
}
func printNum(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	printNum(n - 1)
}
