package pointer

import "fmt"

func Pointer_and_slice() {
	slicedata := []int{1, 2, 3, 4, 5}
	slicedatapointer := &slicedata
	fmt.Println(*slicedatapointer)
	slicedata1 := []string{"a", "b", "c", "d"}
	a := Another_function_slice(slicedata1)
	fmt.Println(a)

}
func Another_function_slice(a []string) []string {
	fmt.Println(a)
	a[0] = "e"
	a[1] = "f"
	a[2] = "g"
	a[3] = "h"
	return a
}
