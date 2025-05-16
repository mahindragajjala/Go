package main

import (
	"array/techniques/twopointertechnique"
)

func main() {
	// array := [5]int{1, 2, 3, 4, 5}
	// array1 := []int{1, 2, 3, 4, 5}
	array1 := []int{6, 3, 0, 4, 0, 2, 1}

	// array2 := []int{1, 2, 3, 4, 5}
	twopointertechnique.MoveZerosToZero(array1)
	twopointertechnique.MoveZerosToEnd(array1)
	// fmt.Println(output)

}
