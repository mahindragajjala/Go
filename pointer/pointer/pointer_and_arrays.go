package pointer

import "fmt"

func Pointer_and_arrays() {
	//Pointers can point to an array, but slices are preferred in Go.
	arr := [3]int{1, 2, 3}
	p := &arr
	fmt.Println((*p)[0])

}
