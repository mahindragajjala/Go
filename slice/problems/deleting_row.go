package problems

import "fmt"

func Deleting_row_slice() {
	slicedata := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	New_slice := append(slicedata[:1], slicedata[2:]...)
	fmt.Println(New_slice)
}
