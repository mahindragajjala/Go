package problems

import "fmt"

func Appending_New_row() {
	data := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	data = append(data, []int{10, 11, 12})
	for _, key := range data {
		fmt.Println(key)
	}

}
