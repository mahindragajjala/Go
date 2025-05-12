package headrecursion

import "fmt"

/*
Print elements of an array from start to end using head recursion.
array := []int{34,25,56,98,32}
array[0]
array[1]
array[2]
array[3]
array[4]
*/
// Head recursion: print from start to end
func Printing_array_Elements(arr []int, index int) {
	if index == len(arr) {
		return // base case: end of array
	}
	// recursive call first (head recursion)
	Printing_array_Elements(arr, index+1)
	// print after recursive call returns
	fmt.Println(arr[index])
}

// Driver function
func Printing_array_Elements_inputs() {
	array := []int{34, 25, 56, 98, 32}
	Printing_array_Elements(array, 0)
}
