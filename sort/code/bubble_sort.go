package code

import "fmt"

func Bubble_sort() {
	array := []int{5, 1, 4, 2, 8}
	n := len(array)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	fmt.Println(array)
}
