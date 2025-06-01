package code

import "fmt"

func checkPositive(a, _ int, z float32) bool {
	return a > 0 && z > 0
}
func Empty_parameter() {
	fmt.Println(checkPositive(5, 100, 3.14)) // Output: true
}
