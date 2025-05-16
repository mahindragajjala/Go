package twopointertechnique

import "fmt"

/*
6. Move Zeroes to End
   Move all 0s to the end of an array while maintaining the relative order of the non-zero elements.

   array := 6,3,0,4,0,2,1
*/
func MoveZerosToZero(array []int) {
	i := 0
	j := len(array) - 1
	for i < j {
		if array[i] == 0 {
			array[i], array[j] = array[j], array[i]
			i++
			j--
		} else {
			i++
		}
	}
	fmt.Println(array)
}
func MoveZerosToEnd(array []int) {
	lastNonZero := 0
	for i := 0; i < len(array); i++ {
		if array[i] != 0 {
			array[lastNonZero], array[i] = array[i], array[lastNonZero]
			lastNonZero++
		}
	}
	fmt.Println(array)
}
