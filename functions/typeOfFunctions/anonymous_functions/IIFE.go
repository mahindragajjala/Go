package anonymousfunctions

import "fmt"

/*
Immediately Invoked Function Expression (IIFE)
*/
func IIFE() {
	result := func(a, b int) int {
		return a + b
	}(4, 5)
	fmt.Println(result) // 9
}
