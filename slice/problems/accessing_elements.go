package problems

import "fmt"

/*
Write a function that returns the element at any given row and column from a [][]int slice.
*/
func Accessing_elements_main() {
	Multiple_array := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	row := 2
	col := 2
	Accessing_elements(Multiple_array, row, col)
}
func Accessing_elements(array [][]int, row int, col int) {
	for _, value := range array {
		fmt.Println(value)
	}

	fmt.Println("Accessing_element:", array[row][col])
}
