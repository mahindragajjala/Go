package anonymousfunctions

import "fmt"

func Passing_as_Argument_to_Another_Function() {
	result := operate(10, 5, func(x, y int) int {
		return x - y
	})
	fmt.Println(result) // 5

}
func operate(a, b int, op func(int, int) int) int {
	return op(a, b)
}
