package anonymousfunctions

import "fmt"

func Assigning_to_a_variable() {
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(2, 3)) // 5

}
