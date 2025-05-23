package rangekeyword

import "fmt"

/*
the range keyword is used in loops to iterate over
elements in arrays, slices, maps, strings, and channels.
*/
func RangeKeyword() {
	Arraydata := [4]int{1, 2, 3, 4}
	slicedata := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for data := range Arraydata {
		fmt.Println("Arraydata:", data)
	}
	for data := range slicedata {
		fmt.Println("slicedata:", data)
	}
}
